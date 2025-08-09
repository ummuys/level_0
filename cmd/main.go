package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/ummuys/level_0/handlers"
	"github.com/ummuys/level_0/logger"
	"github.com/ummuys/level_0/repository"
	"github.com/ummuys/level_0/server"
	"github.com/ummuys/level_0/service"
)

func main() {
	//---ENV
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	//---LOGGER
	logger, err := logger.InitLogger("logs/")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info().Msg("start the program...")
	logger.Info().Msg("env file is loaded")

	db, err := repository.NewDatabase(logger)
	if err != nil {
		logger.Err(err)
		log.Fatal(err)
	}
	logger.Info().Msg("databases initialized")

	//---SERVICES
	orderService := service.NewOrderService(db, logger)
	logger.Info().Msg("service initialized")

	//---HANDLERS
	orderHandler := handlers.NewOrderHandler(orderService, logger)
	logger.Info().Msg("handlers initialized")

	//---CHANNEL FOR GRACEFUL SHUTDOWN
	chSD := make(chan os.Signal, 1)
	signal.Notify(chSD, syscall.SIGINT, syscall.SIGTERM)

	srv := server.InitServer(orderHandler)
	go server.RunServer(srv, logger)

	//---WAIT SHUTDOWN SIGNAL
	<-chSD
	logger.Info().Msg("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//---SHOTDOWN
	srvErr := srv.Shutdown(ctx)
	if srvErr != nil {
		logger.Fatal().Err(srvErr).Msg("server shutdown error")
	}

	dbErr := db.Close()
	if dbErr != nil {
		logger.Fatal().Err(dbErr).Msg("database close error")
	}

	if dbErr != nil || srvErr != nil {
		logger.Fatal().Msg("shutdown with errors")
	} else {
		logger.Info().Msg("server shutdown gracefully")
	}

}
