package main

import (
	"context"
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

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "host=localhost port=5432 user=group_manager password=supersecret dbname=group_manager sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	idProvider := snowflake.NewIdProvider(1)
	peopleRepo := postgres.NewPeopleRepo(conn)
	peopleApi := people.New(idProvider, peopleRepo)
	groupApi := group.New()

	handler := http.NewHander(peopleApi, groupApi)
	router := http.NewRouter(handler)
	server := http.NewServer(":50000", router)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go server.Start()

	<-quit

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("forced shutdown")
	}
}
