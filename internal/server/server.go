package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/handlers"
)

func InitServer(oh handlers.OrderHandler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(HealthEndpoint, oh.Health)

	srv := &http.Server{
		Addr:              ":" + os.Getenv("APP_PORT"),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second, // TODO: настроить время тайм-аутов и мб подкинуть еще
	}
	return srv
}

func RunServer(ctx context.Context, srv *http.Server, logger *zerolog.Logger) error {

	go func() {
		<-ctx.Done()
		offSrv, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		_ = srv.Shutdown(offSrv)
		logger.Info().Msg("close server routine")
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil

}
