package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"ueransim-api/pkg/logger"
)

type Server struct {
	httpServer *http.Server
}

func New(port int, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  1 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	logger.Info("running on", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
