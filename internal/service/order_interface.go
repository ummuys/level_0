package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/repository"
)

type OrderService interface {
}

type orderService struct {
	db     repository.Database
	logger *zerolog.Logger
}
