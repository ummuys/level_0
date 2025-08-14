package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/models"
)

type OrderDB interface {
	Create(pCtx context.Context, orderRawData []byte) error
	Get(pCtx context.Context, oUID string) (models.OrderDataDb, error)
	GetN(pCtx context.Context, n int) ([]models.OrderDataDb, error)
	Close() error
}

type odbPg struct {
	conn   *pgx.Conn
	logger *zerolog.Logger
}
