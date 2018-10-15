// Copyright 2018 James Churchard. All rights reserved.
// Use of this source code is governed by MIT license,
// a copy can be found in the LICENSE file.

// Package sse creates a Server-Sent Events "spec" server for the
// transmission of events to connected clients.
package sse

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// writeTimeout is how long to attempt to write to a client
// before deciding they're too slow and leaving them to fend for themselves.
const writeTimeout time.Duration = time.Second * 1

// client represents a client connection and accepts serialized events.
type client = chan []byte

// Event represents a single server-sent event.
// Data will be processed by the JSON marshaller before being sent.
type Event struct {
	Event string
	Data  interface{}
}

// Server receives events and broadcasts them to all connected clients.
type Server struct {
	// clients stores active connection channels.
	clients map[client]bool

	// connect accepts new clients to add to the clients pool.
	connect chan client

	// disconnect accepts clients to remove from the clients pool.
	disconnect chan client

	// SendEvent accepts an Event{} to broadcast to clients.
	SendEvent chan *Event
}

// New creates a new SSE server, with all required events chans, and
// starts the client management and events dispatch listener loops.
func New() *Server {
	s := &Server{
		clients:    make(map[client]bool),
		connect:    make(chan client),
		disconnect: make(chan client),
		SendEvent:  make(chan *Event, 1),
	}

	go s.run()

	return s
}

// ServeHTTP satifies the http.Handler interface, allowing sse.Server to
// be passed directly as a handler for http-compliant muxes (muxxi?).
// It handles the connection set-up for event-streams, and client pool
// management.
func (s Server) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming is not supported!", http.StatusNotImplemented)

		return
	}

	closeNotifier, ok := w.(http.CloseNotifier)
	if !ok {
		http.Error(w, "Close notification is not supported!", http.StatusNotImplemented)

		return
	}
	closeNotify := closeNotifier.CloseNotify()

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
	flusher.Flush()

	c := make(client)
	s.connect <- c

	for {
		select {
		case <-closeNotify:
			s.disconnect <- c

			return
		case event := <-c:
			w.Write(event)
			flusher.Flush()
		}
	}
}

// run handles the addition/removal of clients, and the broadcast
// of events.
func (s *Server) run() {
	for {
		select {
		case c := <-s.connect:
			s.clients[c] = true
		case c := <-s.disconnect:
			delete(s.clients, c)
		case event := <-s.SendEvent:
			payload := bytes.Buffer{}

			if len(event.Event) > 0 {
				payload.WriteString("event: ")
				payload.WriteString(event.Event)
				payload.WriteByte('\n')
			}

			data, err := json.Marshal(event.Data)
			if err != nil {
				// Silent failure
				return
			}

			payload.WriteString("data: ")
			payload.Write(data)
			payload.Write([]byte{'\n', '\n'})

			for c := range s.clients {
				select {
				case c <- payload.Bytes():
				case <-time.After(writeTimeout):
				}
			}
		}
	}
}
