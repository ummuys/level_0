package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/service"
)

type OrderHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	ordServ service.OrderService
	logger  *zerolog.Logger
}
