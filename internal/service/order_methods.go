package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/repository"
)

func NewOrderService(db repository.OrderDB, logger *zerolog.Logger) OrderService {
	return &orderService{
		db:     db,
		logger: logger,
	}
}

func (ordS orderService) Create(pCtx context.Context, orderRawData []byte) error {
	err := ordS.db.Create(pCtx, orderRawData)
	if err != nil {
		ordS.logger.Error().
			Err(err).
			Msg("db err")
	}
	return err
}
