package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/orewaee/group-manager/internal/config"
	"github.com/rs/zerolog/log"
)

type Server struct {
	httpServer *http.Server
	config     *config.Config
}

func NewServer(cfg *config.Config, router http.Handler) *Server {
	addr := fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)

	return &Server{
		config: cfg,
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  time.Duration(cfg.Http.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.Http.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(cfg.Http.IdleTimeout) * time.Second,
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
