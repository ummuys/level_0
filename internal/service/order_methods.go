package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/cache"
	"github.com/ummuys/level_0/internal/models"
	"github.com/ummuys/level_0/internal/repository"
	"github.com/ummuys/level_0/internal/validation"
)

func NewOrderService(db repository.OrderDB, cache cache.OrderCache, logger *zerolog.Logger) OrderService {
	return &orderService{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

func (ordS *orderService) Create(pCtx context.Context, order models.OrderData) error {

	UID, err := validation.Validate(order)
	if err != nil {
		return fmt.Errorf("validate err: %w", err)
	}

	err = ordS.db.Create(pCtx, order)
	if err != nil {
		ordS.logger.Error().
			Err(err).
			Msg("db err")
		return err
	}

	ordS.cache.Set(UID, order)

	return nil
}

func (ordS *orderService) Get(pCtx context.Context, key string) (models.OrderData, error) {

	cacheInfo := ordS.cache.Get(pCtx, key)

	if cacheInfo.OrderUID != "" {
		return cacheInfo, nil
	}

	dbInfo, err := ordS.db.Get(pCtx, key)
	if err != nil {
		ordS.logger.Error().
			Err(err).
			Msg("db err")
		return models.OrderData{}, err
	}

	ordS.cache.Set(key, dbInfo)

	return dbInfo, nil
}
