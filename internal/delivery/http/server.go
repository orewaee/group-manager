package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 60 * time.Second
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string, router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
	}
}

func (s *Server) Start() error {
	log.Info().Str("addr", s.httpServer.Addr).Msg("start http server")

	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Info().Msg("shutdown http server")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
