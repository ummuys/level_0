package handlers

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/service"
)

func NewOrderHandler(pCtx context.Context, orderService service.OrderService, logger *zerolog.Logger) OrderHandler {
	return &orderHandler{
		pCtx:    pCtx,
		ordServ: orderService,
		logger:  logger,
	}
}

// TODO: WRITE THIS vvvvvvvvv
func (oh *orderHandler) Get(w http.ResponseWriter, r *http.Request) {
	order_uid := r.PathValue("order_uid")
	info, err := oh.ordServ.Get(oh.pCtx, order_uid)
	if err != nil || info.CustomerID == "" {
		w.Write([]byte("problema!"))
	}
	w.Write(info)
}
