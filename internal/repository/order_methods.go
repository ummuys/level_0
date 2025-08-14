package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/models"
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

func (odb *odbPg) Create(pCtx context.Context, orderData models.OrderData) (err error) {
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

	b, n := createBatch(&orderData)
	bSend := tx.SendBatch(txCtx, b)
	defer bSend.Close()

	for i := 0; i < n; i++ {
		_, err = bSend.Exec()
		if err != nil {
			err = fmt.Errorf("batch err: %w", err)
			return
		}
	}

	if err = txCtx.Err(); err != nil {
		err = fmt.Errorf("tx deadline: %w", err)
		return
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

func (odb *odbPg) Get(pCtx context.Context, oUID string) (models.OrderData, error) {
	odb.logger.Debug().Msg("call Get method in OrderDB")

	ctx, cancel := context.WithTimeout(pCtx, time.Second*5)
	defer cancel()

	query := `
	SELECT * FROM orders.info
	JOIN orders.payments ON info.order_uid = payments.order_uid
	JOIN orders.deliveries ON info.order_uid = deliveries.order_uid
	JOIN orders.items ON info.order_uid = items.order_uid
	WHERE order_uid = $1;
	`

	orderData := models.OrderData{}
	err := odb.conn.QueryRow(ctx, query, oUID).Scan(&orderData)
	if err != nil {
		return models.OrderData{}, fmt.Errorf("can't exec a query row: %w", err)
	}

	return orderData, nil
}
