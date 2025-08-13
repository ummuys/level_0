package handlers

import (
	"net/http"

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

}
