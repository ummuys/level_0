package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/repository"
)

func NewOrderService(db repository.Database, logger *zerolog.Logger) OrderService {
	return &orderService{
		db:     db,
		logger: logger,
	}
}
