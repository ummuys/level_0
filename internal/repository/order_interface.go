package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/models"
)

type OrderDB interface {
	Create(pCtx context.Context, orderData models.OrderData) error
	Get(pCtx context.Context, oUID string) (models.OrderData, error)
	Close() error
}

type odbPg struct {
	conn   *pgx.Conn
	logger *zerolog.Logger
}
