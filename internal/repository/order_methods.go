package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/models"
)

func NewOrderDatabase(logger *zerolog.Logger) (OrderDB, error) {

	logger.Info().
		Str("evt", "db.ready.wait").
		Msg("")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var (
		conn *pgx.Conn
		err  error
	)

	for i := 0; i < 5; i++ {
		conn, err = pgx.Connect(ctx, os.Getenv("DB_LINK"))
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
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
	odb.logger.Debug().
		Str("evt", "db.create").
		Msg("")

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

	query := `select orders.insert_order($1::jsonb, $2::timestamptz)`

	_, err = tx.Exec(txCtx, query, orderRawData, time.Now())
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

func (odb *odbPg) Get(pCtx context.Context, oUID string) (models.OrderDataDb, error) {
	odb.logger.Debug().
		Str("evt", "db.get").
		Msg("")

	ctx, cancel := context.WithTimeout(pCtx, time.Second*5)
	defer cancel()

	query, args := createQueryGetN(oUID, 0)

	o := models.OrderDataDb{}
	err := pgxscan.Get(ctx, odb.conn, &o, query, args...)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return models.OrderDataDb{}, fmt.Errorf("can't exec a query row: %w", err)
	}

	items, err := odb.scanItems(ctx, o.OrderUID)
	if err != nil {
		return models.OrderDataDb{}, err
	}
	o.Items = append(o.Items, items...)

	return o, nil
}

func (odb *odbPg) GetN(pCtx context.Context, n int) ([]models.OrderDataDb, error) {
	odb.logger.Debug().
		Str("evt", "db.getN").
		Msg("")

	ctx, cancel := context.WithTimeout(pCtx, time.Second*10)
	defer cancel()

	query, args := createQueryGetN("", n)

	var orders []models.OrderDataDb

	err := pgxscan.Select(ctx, odb.conn, &orders, query, args...)
	if err != nil {
		return nil, err
	}

	for i, o := range orders {
		items, err := odb.scanItems(ctx, o.OrderUID)
		if err != nil {
			return []models.OrderDataDb{}, err
		}
		o.Items = append(o.Items, items...)
		orders[i].Items = o.Items
	}

	return orders, nil
}
