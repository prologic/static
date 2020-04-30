package main

import (
	"net/http"

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
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(
		s.bind,
		s.logger.Handler(
			gziphandler.GzipHandler(
				s.router,
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
		}),
	}

	server.initRoutes()

	return server
}
