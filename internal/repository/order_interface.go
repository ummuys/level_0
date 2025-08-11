package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

type Database interface {
	Close() error
}

type dbPg struct {
	conn   *pgx.Conn
	logger *zerolog.Logger
}
