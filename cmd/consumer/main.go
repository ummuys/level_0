package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// err := godotenv.Load(".env.consumer")
	// if err != nil {
	// 	log.Fatal("no env file: ", err)
	// }

	loggers, err := logger.InitLogger(os.Getenv("LOGS_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	loggers.AppLog.Info().Msg("--------------------------------------------------")
	loggers.AppLog.Info().Msg("-------------- LEVEL 0 BY UMMUYS -----------------")
	loggers.AppLog.Info().Msg("--------------------------------------------------")

	orderDB, err := repository.NewOrderDatabase(loggers.DbLog)
	if err != nil {
		loggers.AppLog.Fatal().
			Err(err).
			Str("evt", "db.ready.fail").
			Msg("")

	}
	loggers.DbLog.Info().
		Str("evt", "db.ready.ok").
		Msg("")

	orderCache, err := cache.NewOrderCache(mainCtx, orderDB, loggers.CchLog)
	if err != nil {
		loggers.AppLog.Fatal().
			Err(err).
			Msg("")
	}

	orderService := service.NewOrderService(mainCtx, orderDB, orderCache, loggers.SvcLog)
	orderHandler := handlers.NewOrderHandler(orderService, loggers.SrvLog)
	serverHandler := handlers.NewServerHandler(loggers.SrvLog)

	srv, err := web.InitServer(serverHandler, orderHandler)
	if err != nil {
		loggers.AppLog.Fatal().
			Err(err).
			Msg("")
	}

	g, ctx := errgroup.WithContext(mainCtx)

	g.Go(func() error {
		loggers.AppLog.Info().Msg("server.start")
		return web.RunServer(ctx, srv, loggers.SrvLog)
	})

	g.Go(func() error {
		loggers.AppLog.Info().Msg("kafka.start")
		return kafka.Kafka(ctx, loggers.KfkLog, orderService)
	})

	gErr := g.Wait()
	if gErr != nil && !errors.Is(gErr, context.Canceled) {
		loggers.AppLog.Error().
			Err(gErr).
			Msg("")
	}

	dbErr := orderDB.Close()
	if dbErr != nil {
		loggers.AppLog.Error().
			Err(dbErr).
			Msg("db.close.fail")
	}

	if dbErr != nil || gErr != nil {
		loggers.AppLog.Error().
			Msg("shutdown.fail")
	} else {
		loggers.AppLog.Info().Msg("shutdown.ok")
	}
}
