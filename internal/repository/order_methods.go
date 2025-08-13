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

func (odb *odbPg) Create(pCtx context.Context, orderRawData []byte) (err error) {
	odb.logger.Debug().Msg("call Create method in OrderDB")

	txCtx, txCancel := context.WithTimeout(pCtx, 7*time.Second)
	defer txCancel()

	var tx pgx.Tx

	tx, err = odb.conn.BeginTx(txCtx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		err = fmt.Errorf("can't begin a tx: %w", err)
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(context.Background())
		}
	}()

	query := `select orders.insert_order($1::jsonb)`

	_, err = tx.Exec(txCtx, query, orderRawData)
	if err != nil {
		err = fmt.Errorf("can't exec a query: %w", err)
		return
	}

	if err = txCtx.Err(); err != nil {
		return fmt.Errorf("tx deadline: %w", err)
	}

	cmtCtx, cmtCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cmtCancel()
	err = tx.Commit(cmtCtx)
	if err != nil {
		err = fmt.Errorf("can't commit a tx: %w", err)
		return
	}

	return nil
}

func (odb *odbPg) Get(pCtx context.Context, oUID string) ([]byte, error) {
	odb.logger.Debug().Msg("call Get method in OrderDB")
	return nil, nil
}
