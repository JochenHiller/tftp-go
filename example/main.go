/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/betawaffle/tftp-go"
)

type Handler struct {
	Path string
}

func (h Handler) ReadFile(c tftp.Conn, filename string) (tftp.ReadCloser, error) {
	log.Printf("Request from %s to read %s", c.RemoteAddr(), filename)
	return os.OpenFile(path.Join(h.Path, filename), os.O_RDONLY, 0)
}

func (h Handler) WriteFile(c tftp.Conn, filename string) (tftp.WriteCloser, error) {
	log.Printf("Request from %s to write %s", c.RemoteAddr(), filename)
	return os.OpenFile(path.Join(h.Path, filename), os.O_WRONLY, 0644)
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: tftp-go <ip-address> [<port>]\n")
		fmt.Fprintf(os.Stderr, "  port defaults to 69 if not specified\n")
		os.Exit(1)
	} else {
		h := Handler{Path: pwd}

		port := "69"
		if len(os.Args) > 2 {
			port = os.Args[2]
		}
		addr := os.Args[1] + ":" + port
		fmt.Printf("Running tftp-server on address %s ...\n", addr)
		err = tftp.ListenAndServe(addr, h)
		panic(err)
	}
}
