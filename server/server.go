package server

import (
	"fmt"
	"net/http"

	"github.com/KonradChlupka/berglen-led/colourutils"
	"github.com/KonradChlupka/berglen-led/engine"
)

type Server interface {
	Serve() error
}

type server struct {
	engine engine.Engine

	port string
}

func NewServer(
	engine engine.Engine,
	opts ...opt,
) *server {
	s := &server{
		engine: engine,
		port:   "8888",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type opt func(*server)

func WithPort(port string) opt {
	return func(s *server) {
		s.port = port
	}
}

// Serve starts the server.
func (s *server) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "sup fam")
	})

	http.HandleFunc("/wipe", func(w http.ResponseWriter, req *http.Request) {
		err := s.engine.ColourWipe(colourutils.BLUE)
		if err != nil {
			fmt.Fprintf(w, "Error while wiping: %s\n", err)
		}
	})

	return http.ListenAndServe(":"+s.port, nil)
}
