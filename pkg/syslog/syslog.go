// Copyright 2018 James Churchard. All rights reserved.
// Use of this source code is governed by MIT license,
// a copy can be found in the LICENSE file.

// Package syslog listens on specified ports/protocols/paths for syslog
// messages (RFC3164, RFC5424, or RFC6587) and emits parsed messages.
package syslog

import (
	slog "gopkg.in/mcuadros/go-syslog.v2"
)

// Server represents a single syslog server, capable of multiple listeners.
type Server struct {
	channel    slog.LogPartsChannel
	server     *slog.Server
	ReceiveLog chan map[string]interface{}
}

// New creates a new Server capable of handling any valid syslog
// format, and creates the log emit chan (ReceiveLog).
func New() *Server {
	s := &Server{
		channel:    make(slog.LogPartsChannel),
		server:     slog.NewServer(),
		ReceiveLog: make(chan map[string]interface{}),
	}

	s.server.SetFormat(slog.Automatic)
	s.server.SetHandler(slog.NewChannelHandler(s.channel))

	return s
}

// Start finalizes configuration and starts listeners.
func (s *Server) Start() error {
	if err := s.server.Boot(); err != nil {
		return err
	}

	go func() {
		for l := range s.channel {
			s.ReceiveLog <- l
		}
	}()

	go s.server.Wait()

	return nil
}

func (s *Server) Shutdown() {
	s.server.Kill()
}

// ListenTCP configures the syslog server to listen for messages
// over TCP at the specified address.
func (s *Server) ListenTCP(addr string) error {
	return s.server.ListenTCP(addr)
}

// ListenUDP configures the syslog server to listen for messages
// over UDP at the specified address.
func (s *Server) ListenUDP(addr string) error {
	return s.server.ListenUDP(addr)
}

// ListenUnix configures the syslog server to listen for messages
// on the specified unix socket path.
func (s *Server) ListenUnix(path string) error {
	return s.server.ListenUnixgram(path)
}
