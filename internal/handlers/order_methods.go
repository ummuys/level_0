package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/service"
)

func NewOrderHandler(orderService service.OrderService, logger *zerolog.Logger) OrderHandler {
	return &orderHandler{
		orderService: orderService,
		logger:       logger,
	}
}

func (or *orderHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
