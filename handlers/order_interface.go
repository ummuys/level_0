package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/service"
)

type OrderHandler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	orderService service.OrderService
	logger       *zerolog.Logger
}
