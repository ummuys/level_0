package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/repository"
)

func NewOrderService(db repository.Database, logger *zerolog.Logger) OrderService {
	return &orderService{
		db:     db,
		logger: logger,
	}
}
