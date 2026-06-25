package http

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(addr string, router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Info().Str("addr", s.httpServer.Addr).Msg("start http server")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Info().Msg("shutdown http server")
	return s.httpServer.Shutdown(ctx)
}
