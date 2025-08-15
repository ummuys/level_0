package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/cache"
	"github.com/ummuys/level_0/internal/models"
	"github.com/ummuys/level_0/internal/repository"
	"github.com/ummuys/level_0/internal/validation"
)

func NewOrderService(pCtx context.Context, db repository.OrderDB, cache cache.OrderCache, logger *zerolog.Logger) OrderService {
	return &orderService{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

func (ordS *orderService) Create(pCtx context.Context, orderRawData []byte, order models.OrderData) error {

	UID, err := validation.Validate(order)
	if err != nil {
		return fmt.Errorf("validate err: %w", err)
	}

	err = ordS.db.Create(pCtx, orderRawData)
	if err != nil {
		return err
	}

	ordS.cache.Set(UID, orderRawData)

	return nil
}

func (ordS *orderService) Get(pCtx context.Context, key string) ([]byte, error) {

	cacheInfo := ordS.cache.Get(pCtx, key)

	if cacheInfo != nil {
		return cacheInfo, nil
	}

	dbInfo, err := ordS.db.Get(pCtx, key)
	if err != nil {
		return nil, err
	}

	if dbInfo.OrderUID == "" {
		return nil, nil
	}

	//OrderDataDb -> OrderData
	order := validation.Convert(dbInfo)

	b, _ := json.Marshal(order)

	ordS.cache.Set(key, b)

	return b, nil
}
