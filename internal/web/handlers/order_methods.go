package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/models"
	"github.com/ummuys/level_0/internal/service"
)

func NewOrderHandler(orderService service.OrderService, logger *zerolog.Logger) OrderHandler {
	return &orderHandler{
		ordServ: orderService,
		logger:  logger,
	}
}

// GetOrder godoc
// @Summary Получить заказ
// @Description Получить заказ по order_uid
// @Tags orders
// @Accept json
// @Produce json
// @Param order_uid path string true "UID заказа"
// @Success 200 {object} models.OrderData
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/order/{order_uid} [get]
func (oh *orderHandler) Get(w http.ResponseWriter, r *http.Request) {
	oh.logger.Debug().
		Str("evt", "handler.Get").
		Msg("")

	w.Header().Set("Content-Type", "application/json")

	order_uid := r.PathValue("order_uid")

	if order_uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		oh.logger.Debug().
			Str("evt", "handler.get.fail").
			Msg("")

		if err := json.NewEncoder(w).Encode(models.ErrorResponse{Error: "order_uid required "}); err != nil {
			http.Error(w, "err encode JSON", http.StatusInternalServerError)
		}
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	info, err := oh.ordServ.Get(ctx, order_uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(models.ErrorResponse{Error: "internal err"}); err != nil {
			http.Error(w, "err encode JSON", http.StatusInternalServerError)
		}
		oh.logger.Debug().
			Str("evt", "handler.get.fail").
			Msg("")
		return
	}

	if info == nil {
		w.WriteHeader(http.StatusNotFound)

		oh.logger.Warn().
			Str("order_uid", order_uid).
			Msg("order didn't found")

		oh.logger.Debug().
			Str("evt", "handler.get.fail").
			Msg("")

		if err := json.NewEncoder(w).Encode(models.ErrorResponse{Error: "order didn't found"}); err != nil {
			http.Error(w, "err encode JSON", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(info)
	oh.logger.Info().
		Str("order_uid", order_uid).
		Msg("order founds")
	oh.logger.Debug().
		Str("evt", "handler.get.ok").
		Msg("")
}
