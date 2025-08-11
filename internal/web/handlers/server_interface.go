package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
)

type ServerHandler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type serverHandler struct {
	logger *zerolog.Logger
}
