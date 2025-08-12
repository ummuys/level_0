package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

func NewOrderDatabase(logger *zerolog.Logger) (OrderDB, error) {

	//TODO:: change CTX
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
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

	return &odbPg{
		conn:   conn,
		logger: logger,
	}, nil

}

func (odb *odbPg) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := odb.conn.Close(ctx); err != nil {
		return fmt.Errorf("can't close db conn: %v", err)
	}
	return nil
}

func (odb *odbPg) Create(pCtx context.Context, orderRawData []byte) error {

	ctx, cancel := context.WithTimeout(pCtx, 5*time.Second)
	defer cancel()

	tx, err := odb.conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `select orders.insert_order($1::jsonb)`

	if _, err := tx.Exec(ctx, query, orderRawData); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (odb *odbPg) Get(pCtx context.Context, oUID string) error {
	return nil
}
