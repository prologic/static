package main

import (
	"log"
	"net/http"

	// Logging
	"github.com/unrolled/logger"

	"github.com/NYTimes/gziphandler"
	"github.com/julienschmidt/httprouter"
)

// Server ...
type Server struct {
	bind   string
	root   string
	router *httprouter.Router

	// Logger
	logger *logger.Logger
}

// ListenAndServe ...
func (s *Server) ListenAndServe() {
	log.Fatal(
		http.ListenAndServe(
			s.bind,
			s.logger.Handler(
				gziphandler.GzipHandler(
					s.router,
				),
			),
		),
	)
}

func (s *Server) initRoutes() {
	s.router.ServeFiles("/*filepath", http.Dir(s.root))
}

// NewServer ...
func NewServer(bind, root string) *Server {
	server := &Server{
		bind:   bind,
		root:   root,
		router: httprouter.New(),

		// Logger
		logger: logger.New(logger.Options{
			Prefix:               "static",
			RemoteAddressHeaders: []string{"X-Forwarded-For"},
			OutputFlags:          log.LstdFlags,
		}),
	}

	server.initRoutes()

	return server
}
