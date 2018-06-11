// Copyright (C) 2017 Michał Matczuk
// Use of this source code is governed by an AGPL-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

const usage1 string = `Usage: tunnel

`

const usage2 string = `Config:
config.yaml:
	server_addr: gerbil-jugaad-tunneld.service.domain.tld:2534
	tunnels:
	  gerbil:
	    proto: http
	    addr: localhost:8080
	    host: jugaad.domain.tld

Author:
	Written by M. Matczuk (mmatczuk@gmail.com)

Patched by:
	Gautam (g@ut.am)

Bugs:
	Submit bugs for patched version to g@ut.am
`

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}
}

type options struct {
	config   string
	logLevel int
	version  bool
	command  string
	args     []string
}

func parseArgs() (*options, error) {
	logLevel := flag.Int("log-level", 3, "Level of messages to log, 0-3")
	flag.Parse()

	opts := &options{
		config:   "/app/gerbil/client.yaml",
		logLevel: *logLevel,
		version:  false,
		command:  "start",
	}
	opts.args = []string{"gerbil"}
	return opts, nil
}
