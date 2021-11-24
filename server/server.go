package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/KonradChlupka/berglen-led/colourutils"
	"github.com/KonradChlupka/berglen-led/engine"
)

type Server interface {
	Serve() error
}

type server struct {
	engine engine.Engine

	mu               *sync.RWMutex
	globalProgram    engine.LEDProgram
	temporaryProgram engine.TemporaryLEDProgram

	port string
}

func NewServer(
	engine engine.Engine,
	opts ...opt,
) *server {
	s := &server{
		engine: engine,
		port:   "8888",

		mu: &sync.RWMutex{},
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

func WithProgram(program engine.LEDProgram) opt {
	return func(s *server) {
		s.globalProgram = program
	}
}

func (s *server) globalRunner() {
	var err error
	for err == nil {
		s.mu.RLock()

		if s.temporaryProgram != nil {
			for !s.temporaryProgram.IsDone() {
				err = s.temporaryProgram.RenderFrame()
				if err != nil {
					break
				}
			}
			s.temporaryProgram = nil
		} else {
			err = s.globalProgram.RenderFrame()
		}

		s.mu.RUnlock()
	}
	fmt.Printf("Global Runner exited!: %v", err)
}

// Serve starts the server.
func (s *server) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "sup fam")
	})

	http.HandleFunc("/wipe", func(w http.ResponseWriter, req *http.Request) {
		// Attempt to get access to the LED's.
		s.mu.Lock()
		defer s.mu.Unlock()

		colourWipe, err := s.engine.ColourWipe(colourutils.BLUE)
		if err != nil {
			fmt.Fprintf(w, "Error while wiping: %s\n", err)
		}
		s.temporaryProgram = colourWipe
	})

	// Start up global runner.
	go s.globalRunner()

	return http.ListenAndServe(":"+s.port, nil)
}
