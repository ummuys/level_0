package http

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	config "github.com/ummuys/level_0/internal/config/server"
	"github.com/ummuys/level_0/internal/web/handlers"

	_ "github.com/ummuys/level_0/docs"
)

func InitServer(sh handlers.ServerHandler, oh handlers.OrderHandler) (*http.Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc(HealthEndpoint, sh.Health)
	mux.HandleFunc(GetOrderEndpoint, oh.Get)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	port, err := config.ParseServerEnv()
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second, // TODO: настроить время тайм-аутов и мб подкинуть еще
	}
	return srv, nil
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
