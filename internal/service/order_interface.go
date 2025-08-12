package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/repository"
)

type OrderService interface {
	Create(pCtx context.Context, orderRawData []byte) error
}

type orderService struct {
	db     repository.OrderDB
	logger *zerolog.Logger
}
