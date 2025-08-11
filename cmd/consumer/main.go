package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ummuys/level_0/internal/handlers"
	"github.com/ummuys/level_0/internal/kafka"
	"github.com/ummuys/level_0/internal/logger"
	"github.com/ummuys/level_0/internal/repository"
	"github.com/ummuys/level_0/internal/server"
	"github.com/ummuys/level_0/internal/service"
	"golang.org/x/sync/errgroup"
)

func main() {

	mainCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := godotenv.Load(".env.consumer")
	if err != nil {
		log.Fatal("no env file: ", err)
	}

	baseLog, kfkLog, srvLog, err := logger.InitLogger(os.Getenv("LOGS_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	baseLog.Info().Msg("--------------------------------------------------")
	baseLog.Info().Msg("-------------- LEVEL 0 BY UMMUYS -----------------")
	baseLog.Info().Msg("--------------------------------------------------")

	db, err := repository.NewDatabase(baseLog)
	if err != nil {
		baseLog.Fatal().
			Err(err).
			Msg("")

	}
	baseLog.Info().Msg("database initialized")

	orderService := service.NewOrderService(db, baseLog)
	orderHandler := handlers.NewOrderHandler(orderService, baseLog)

	srv := server.InitServer(orderHandler)

	g, ctx := errgroup.WithContext(mainCtx)

	g.Go(func() error {
		baseLog.Info().Msg("start the server")
		return server.RunServer(ctx, srv, srvLog)
	})

	g.Go(func() error {
		baseLog.Info().Msg("start the kafka")
		return kafka.Kafka(ctx, kfkLog)
	})

	gErr := g.Wait()
	if gErr != nil && !errors.Is(gErr, context.Canceled) {
		baseLog.Error().
			Err(gErr).
			Msg("")
	}

	dbErr := db.Close()
	if dbErr != nil {
		baseLog.Error().
			Err(dbErr).
			Msg("database close error")
	}

	if dbErr != nil || gErr != nil {
		baseLog.Error().
			Msg("shutdown with errors")
	} else {
		baseLog.Info().Msg("server shutdown gracefully")
	}
}
