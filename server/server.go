package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/handlers"
)

func InitServer(oh handlers.OrderHandler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(HealthEndpoint, oh.Health)

	srv := &http.Server{
		Addr:              os.Getenv("APP_PORT"),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second, // TODO: настроить время тайм-аутов и мб подкинуть еще
	}
	return srv
}

func RunServer(srv *http.Server, logger *zerolog.Logger) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		msg := fmt.Sprintf("listen error: %v", err)
		logger.Fatal().Msg(msg)
	}
}
