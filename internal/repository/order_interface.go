package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

type OrderDB interface {
	Create(pCtx context.Context, orderRawData []byte) error
	Get(pCtx context.Context, oUID string) error
	Close() error
}

type odbPg struct {
	conn   *pgx.Conn
	logger *zerolog.Logger
}
