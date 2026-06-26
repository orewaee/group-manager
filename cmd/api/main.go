package main

import (
	"context"
	"errors"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/orewaee/group-manager/internal/delivery/http"
	"github.com/orewaee/group-manager/internal/infra/postgres"
	"github.com/orewaee/group-manager/internal/infra/snowflake"
	"github.com/orewaee/group-manager/internal/usecase/group"
	"github.com/orewaee/group-manager/internal/usecase/people"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const shutdownTimeout = 30 * time.Second

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx := context.Background()

	conn, err := pgx.Connect(ctx,
		"host=localhost port=5432 user=group_manager password=supersecret dbname=group_manager sslmode=disable",
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := conn.Close(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to close db connection")
		}
	}()

	idProvider := snowflake.NewIdProvider(1)
	peopleRepo := postgres.NewPeopleRepo(conn)
	groupRepo := postgres.NewGroupRepo(conn)
	peopleApi := people.New(idProvider, peopleRepo)
	groupApi := group.New(idProvider, groupRepo, peopleRepo)

	handler := http.NewHander(peopleApi, groupApi)
	router := http.NewRouter(handler)
	server := http.NewServer(":50000", router)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.Start()
		if err != nil && !errors.Is(err, stdhttp.ErrServerClosed) {
			log.Error().Err(err).Msg("server failed")
			os.Exit(1)
		}
	}()

	<-quit

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, shutdownTimeout)

	errShutdown := server.Shutdown(shutdownCtx)

	shutdownCancel()

	if errShutdown != nil {
		log.Error().Err(errShutdown).Msg("forced shutdown")
	}
}
