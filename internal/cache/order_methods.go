package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/models"
	"github.com/ummuys/level_0/internal/repository"
)

type orderNTime struct {
	info   models.OrderData
	expire time.Time
}

type orderCache struct {
	mu       sync.RWMutex
	m        map[string]orderNTime
	db       repository.OrderDB
	cap      int
	ttlOrder time.Duration
	logger   *zerolog.Logger
}

func NewOrderCache(db repository.OrderDB, chcLog *zerolog.Logger) (OrderCache, error) {

	capStr := os.Getenv("CACHE_CAPACITY")
	cap, err := strconv.Atoi(capStr)
	if err != nil {
		return nil, fmt.Errorf("cache capacity err: %w", err)
	}

	ttlOrderSecStr := os.Getenv("TTL_ORDER")
	ttlOrderSec, err := strconv.Atoi(ttlOrderSecStr)
	if err != nil {
		return nil, fmt.Errorf("cache capacity err: %w", err)
	}

	return &orderCache{
		m:        make(map[string]orderNTime),
		db:       db,
		cap:      cap,
		ttlOrder: time.Second * time.Duration(ttlOrderSec),
		logger:   chcLog,
	}, nil
}

func (oc *orderCache) Set(orderUID string, orderData models.OrderData) {
	oc.logger.Debug().Msg("call Set")
	expire := time.Time{}

	if oc.ttlOrder > 0 {
		expire = time.Now().Add(oc.ttlOrder)
	}

	// Lock - we fill the map
	oc.mu.Lock()
	defer oc.mu.Unlock()

	if len(oc.m) > oc.cap {
		oc.callClearExprired()
		for len(oc.m) >= oc.cap {
			for k := range oc.m {
				delete(oc.m, k)
				break
			}
		}
	}

	oc.m[orderUID] = orderNTime{
		info:   orderData,
		expire: expire,
	}

}

func (oc *orderCache) Get(pCtx context.Context, orderUID string) models.OrderData {
	oc.logger.Debug().Msg("call Get")

	oc.mu.RLock()
	defer oc.mu.RUnlock()

	item, ok := oc.m[orderUID]
	if !ok {
		return models.OrderData{}
	}

	return item.info
}

func (oc *orderCache) callClearExprired() {
	oc.logger.Debug().Msg("call CallClearExprired")
	oc.mu.Lock()
	oc.clearExpired()
	oc.mu.Unlock()
}

func (oc *orderCache) clearExpired() {
	oc.logger.Debug().Msg("call clearExprired")
	// ttlOrder == 0 --> no ttl
	if oc.ttlOrder < 0 {
		return
	}

	// Lock - we delete items in the map
	now := time.Now()
	for key, item := range oc.m {
		if now.After(item.expire) {
			oc.logger.Info().
				Str("key", key).
				Msg("Time expired")
			delete(oc.m, key)
		}
	}
}
