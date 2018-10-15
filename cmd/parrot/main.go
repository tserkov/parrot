// Copyright 2018 James Churchard. All rights reserved.
// Use of this source code is governed by MIT license,
// a copy can be found in the LICENSE file.

// Package main is the Parrot CLI binary entrypoint
package main

import (
	"flag"
	"os"

	"github.com/tserkov/parrot/pkg/broker"
)

// Main assembles the configuration from command-line arguments
// and, if valid, passes it to the broker to start the party.
func main() {
	c := &broker.Config{}

	flag.Var(&c.Forwarders, "forward", "forward received logs to the specified syslog server")
	flag.Var(&c.Listeners, "listen", "listen for logs on the specified location (required)")
	flag.StringVar(&c.Web, "web", "127.0.0.1:8080", "host:port for dashboard webserver")
	flag.BoolVar(&c.Silent, "silent", false, "specify to disable info logging")
	flag.Parse()

	// There must be at least one listener.
	if len(c.Listeners) == 0 {
		flag.PrintDefaults()

		os.Exit(2)
	}

	// Start the party.
	if err := broker.Start(c); err != nil {
		panic(err)
	}
}
