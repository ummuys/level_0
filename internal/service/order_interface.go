package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/cache"
	"github.com/ummuys/level_0/internal/models"
	"github.com/ummuys/level_0/internal/repository"
)

type OrderService interface {
	Create(pCtx context.Context, orderRawData []byte, order models.OrderData) error
	Get(pCtx context.Context, key string) ([]byte, error)
}

type orderService struct {
	db     repository.OrderDB
	cache  cache.OrderCache
	logger *zerolog.Logger
}
