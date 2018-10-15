// Copyright 2018 James Churchard. All rights reserved.
// Use of this source code is governed by MIT license,
// a copy can be found in the LICENSE file.

// Package webserver serves the web dashboard from either embedded files (prod)
// or app directory (dev), as well as provide the SSE server endpoint.
package webserver

import (
	"context"
	"net/http"
	// "path/filepath"
	"time"

	"github.com/tserkov/parrot/pkg/sse"

	"github.com/GeertJohan/go.rice"
)

type Server struct {
	server *http.Server
}

// Start beings serving the web dashboard at the specified address, and
// sets up the endpoint for the provided SSE server.
func New(addr string, sseServer *sse.Server) *Server {
	r := http.NewServeMux()

	r.Handle("/squawk", sseServer)

	// Serves the dashboard app from either the local app directory in dev,
	// or embedded files in prod.
	// dir, _ := filepath.Abs("./app/dist")
	box := rice.MustFindBox("../../app/dist")
	r.Handle("/", http.FileServer(box.HTTPBox()))

	return &Server{
		server: &http.Server{
			Handler: r,
			Addr:    addr,
		},
	}
}

// Start begins listening and serving the dashboard & SSE endpoint.
func (s *Server) Start() {
	go s.server.ListenAndServe()
}

// Shutdown cleanly stops the server listener.
func (s *Server) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	return s.server.Shutdown(ctx)
}
