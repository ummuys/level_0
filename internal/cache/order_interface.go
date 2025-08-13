package cache

import "context"

type OrderCache interface {
	Set(orderUID string, orderInfo []byte)
	Get(pCtx context.Context, orderUID string) []byte
}
