package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

func NewDatabase(logger *zerolog.Logger) (Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var (
		conn *pgx.Conn
		err  error
	)

	for i := 0; i < 10; i++ {
		conn, err = pgx.Connect(ctx, os.Getenv("DB_LINK"))
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("db didn't pinged: %w", err)
	}

	return &dbPg{
		conn:   conn,
		logger: logger,
	}, nil

}

func (db *dbPg) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := db.conn.Close(ctx); err != nil {
		return fmt.Errorf("can't close db conn: %v", err)
	}
	return nil
}
