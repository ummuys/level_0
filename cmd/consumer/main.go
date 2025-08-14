package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ummuys/level_0/internal/cache"
	"github.com/ummuys/level_0/internal/kafka"
	"github.com/ummuys/level_0/internal/logger"
	"github.com/ummuys/level_0/internal/repository"
	"github.com/ummuys/level_0/internal/service"
	web "github.com/ummuys/level_0/internal/web"
	"github.com/ummuys/level_0/internal/web/handlers"
	"golang.org/x/sync/errgroup"
)

func main() {

	mainCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//TODO: fix this
	err := godotenv.Load(".env.consumer")
	if err != nil {
		log.Fatal("no env file: ", err)
	}

	appLog, kfkLog, srvLog, cchLog, err := logger.InitLogger(os.Getenv("LOGS_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	appLog.Info().Msg("--------------------------------------------------")
	appLog.Info().Msg("-------------- LEVEL 0 BY UMMUYS -----------------")
	appLog.Info().Msg("--------------------------------------------------")

	orderDB, err := repository.NewOrderDatabase(appLog)
	if err != nil {
		appLog.Fatal().
			Err(err).
			Msg("")

	}
	appLog.Info().Msg("database initialized")

	orderCache, err := cache.NewOrderCache(orderDB, cchLog)
	if err != nil {
		appLog.Fatal().
			Err(err).
			Msg("")
	}

	orderService := service.NewOrderService(orderDB, orderCache, srvLog)
	orderHandler := handlers.NewOrderHandler(mainCtx, orderService, srvLog)
	serverHandler := handlers.NewServerHandler(srvLog)

	srv := web.InitServer(serverHandler, orderHandler)
	g, ctx := errgroup.WithContext(mainCtx)

	g.Go(func() error {
		appLog.Info().Msg("start the server")
		return web.RunServer(ctx, srv, srvLog)
	})

	g.Go(func() error {
		appLog.Info().Msg("start the kafka")
		return kafka.Kafka(ctx, kfkLog, orderService)
	})

	gErr := g.Wait()
	if gErr != nil && !errors.Is(gErr, context.Canceled) {
		appLog.Error().
			Err(gErr).
			Msg("")
	}

	dbErr := orderDB.Close()
	if dbErr != nil {
		appLog.Error().
			Err(dbErr).
			Msg("database close error")
	}

	if dbErr != nil || gErr != nil {
		appLog.Error().
			Msg("app shutdown with errors")
	} else {
		appLog.Info().Msg("app shutdown gracefully")
	}
}
