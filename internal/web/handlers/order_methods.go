package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/service"
)

func NewOrderHandler(orderService service.OrderService, logger *zerolog.Logger) OrderHandler {
	return &orderHandler{
		ordServ: orderService,
		logger:  logger,
	}
}

// TODO: WRITE THIS vvvvvvvvv
func (oh *orderHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	order_uid := r.PathValue("order_uid")
	if order_uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("order_uid is requared"))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	info, err := oh.ordServ.Get(ctx, order_uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if info == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("order didn't found"))

		oh.logger.Info().
			Str("order_uid", order_uid).
			Msg("order didn't found")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(info)
	oh.logger.Info().
		Str("order_uid", order_uid).
		Msg("order founds")
}
