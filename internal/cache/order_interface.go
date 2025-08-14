package cache

import (
	"context"

	"github.com/ummuys/level_0/internal/models"
)

type OrderCache interface {
	Set(orderUID string, orderData models.OrderData)
	Get(pCtx context.Context, orderUID string) models.OrderData
}
