package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/repository"
)

type orderNTime struct {
	info   []byte
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

func NewOrderCache(pCtx context.Context, db repository.OrderDB, chcLog *zerolog.Logger) (OrderCache, error) {

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

	ordC := orderCache{
		m:        make(map[string]orderNTime),
		db:       db,
		cap:      cap,
		ttlOrder: time.Second * time.Duration(ttlOrderSec),
		logger:   chcLog,
	}

	err = ordC.fillCache(pCtx)
	if err != nil {
		return nil, err
	}

	return &ordC, nil
}

func (ordC *orderCache) Set(orderUID string, orderInfo []byte) {
	ordC.logger.Debug().Msg("call Set")
	expire := time.Time{}

	if ordC.ttlOrder > 0 {
		expire = time.Now().Add(ordC.ttlOrder)
	}

	// Lock - we fill the map
	ordC.mu.Lock()
	defer ordC.mu.Unlock()

	if len(ordC.m) > ordC.cap {
		ordC.callClearExprired()
		for len(ordC.m) >= ordC.cap {
			for k := range ordC.m {
				delete(ordC.m, k)
				break
			}
		}
	}

	ordC.m[orderUID] = orderNTime{
		info:   orderInfo,
		expire: expire,
	}

}

func (ordC *orderCache) Get(pCtx context.Context, orderUID string) []byte {
	ordC.logger.Debug().Msg("call Get")

	ordC.mu.RLock()
	defer ordC.mu.RUnlock()

	item, ok := ordC.m[orderUID]
	if !ok {
		return nil
	}

	return item.info
}

func (ordC *orderCache) callClearExprired() {
	ordC.logger.Debug().Msg("call CallClearExprired")
	ordC.mu.Lock()
	ordC.clearExpired()
	ordC.mu.Unlock()
}

func (ordC *orderCache) clearExpired() {
	ordC.logger.Debug().Msg("call clearExprired")
	// ttlOrder == 0 --> no ttl
	if ordC.ttlOrder < 0 {
		return
	}

	// Lock - we delete items in the map
	now := time.Now()
	for key, item := range ordC.m {
		if now.After(item.expire) {
			ordC.logger.Info().
				Str("key", key).
				Msg("Time expired")
			delete(ordC.m, key)
		}
	}
}

func (ordC *orderCache) fillCache(pCtx context.Context) error {
	ordC.logger.Debug().Msg("call fillCache")

	orders, err := ordC.db.GetN(pCtx, ordC.cap)
	if err != nil {
		return err
	}

	i := 0
	for _, o := range orders {
		oB, err := json.Marshal(o)
		if err != nil {
			return fmt.Errorf("can't unmarshall orders: %w", err)
		}
		ordC.Set(o.OrderUID, oB)
		i++
	}

	if i == 0 {
		ordC.logger.Info().Msg("Nothing to load into db")
	} else {
		ordC.logger.Info().
			Int("Amount", i).
			Msg("Loaded a orders from db")
	}
	return nil
}
