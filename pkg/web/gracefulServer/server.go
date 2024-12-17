package gracefulServer

import (
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	gracefulTimeout time.Duration
}

func NewServer(addr string, handler http.Handler, gracefulTimeout time.Duration) *Server {
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		gracefulTimeout: gracefulTimeout,
	}
}

func (s *Server) ListenAndServe() error {

	return s.server.ListenAndServe()
}
