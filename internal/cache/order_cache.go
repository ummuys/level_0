package cache

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/repository"
)

type OrderNTime struct {
	info   []byte
	expire time.Time
}

type OrderCache struct {
	mu       sync.RWMutex
	m        map[string]OrderNTime
	db       repository.OrderDB
	capacity int
	ttlOrder time.Duration
	logger   *zerolog.Logger
}

func NewOrderCache(db repository.OrderDB, capacity int, ttlOrder time.Duration, chcLog *zerolog.Logger) *OrderCache {
	return &OrderCache{
		m:        make(map[string]OrderNTime),
		db:       db,
		capacity: capacity,
		ttlOrder: ttlOrder,
		logger:   chcLog,
	}
}

func (oc *OrderCache) Set(orderUID string, orderInfo []byte) {
	oc.logger.Debug().Msg("call Set")
	expire := time.Time{}

	if oc.ttlOrder > 0 {
		expire = time.Now().Add(oc.ttlOrder)
	}

	// Lock - we fill the map
	oc.mu.Lock()
	defer oc.mu.Unlock()

	if len(oc.m) > oc.capacity {
		oc.CallClearExprired()
		for len(oc.m) >= oc.capacity {
			for k := range oc.m {
				delete(oc.m, k)
				break
			}
		}
	}

	oc.m[orderUID] = OrderNTime{
		info:   orderInfo,
		expire: expire,
	}

}

func (oc *OrderCache) cacheGet(orderUID string) []byte {
	oc.logger.Debug().Msg("call cacheGet")

	// Read lock - we read the map
	oc.mu.RLock()
	defer oc.mu.RUnlock()

	// Didn't found
	cacheItem, ok := oc.m[orderUID]
	if !ok {
		return nil
	}

	return cacheItem.info

}

func (oc *OrderCache) Get(pCtx context.Context, orderUID string) ([]byte, error) {
	oc.logger.Debug().Msg("call Get")
	// If we found order in cache
	cacheItem := oc.cacheGet(orderUID)
	if cacheItem != nil {
		return cacheItem, nil
	}

	// If we don't, we will check in db
	dbItem, err := oc.db.Get(pCtx, orderUID)
	if err != nil {
		return nil, err
	}

	// Set order in cache
	oc.Set(orderUID, dbItem)
	return dbItem, nil
}

func (oc *OrderCache) CallClearExprired() {
	oc.logger.Debug().Msg("call CallClearExprired")
	oc.mu.Lock()
	oc.clearExpired()
	oc.mu.Unlock()
}

func (oc *OrderCache) clearExpired() {
	oc.logger.Debug().Msg("call clearExprired")
	// ttlOrder == 0 --> no ttl
	if oc.ttlOrder < 0 {
		return
	}

	// Lock - we delete items in the map
	now := time.Now()
	for key, item := range oc.m {
		if now.After(item.expire) {
			delete(oc.m, key)
		}
	}
}
