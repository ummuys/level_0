package repository

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/ummuys/level_0/internal/models"
)

func (odb *odbPg) scanItems(ctx context.Context, orderUID string) ([]models.ItemData, error) {

	var items []models.ItemData
	err := pgxscan.Select(ctx, odb.conn, &items, querySelectItems, orderUID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
