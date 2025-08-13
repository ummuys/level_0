package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
)

func NewServerHandler(logger *zerolog.Logger) ServerHandler {
	return &serverHandler{
		logger: logger,
	}
}

func (sr *serverHandler) Health(w http.ResponseWriter, r *http.Request) {
	sr.logger.Debug().Msg("call Health method")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
