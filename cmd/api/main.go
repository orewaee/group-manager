package main

import (
	"context"
	"log"
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
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=group_manager dbname=group_manager sslmode=disable")
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

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	<-quit

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("forced shutdown:", err)
	}
}
