package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"
	config "github.com/ummuys/level_0/internal/config/cache"
	"github.com/ummuys/level_0/internal/repository"
	"github.com/ummuys/level_0/internal/validation"
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
	cchEnv, err := config.ParseCacheEnv()
	if err != nil {
		return nil, err
	}

	ordC := orderCache{
		m:        make(map[string]orderNTime),
		db:       db,
		cap:      cchEnv.Capacity,
		ttlOrder: cchEnv.TTLOrder,
		logger:   chcLog,
	}

	err = ordC.fillCache(pCtx)
	if err != nil {
		return nil, err
	}

	return &ordC, nil
}

func (ordC *orderCache) Set(orderUID string, orderInfo []byte) {
	ordC.logger.Debug().
		Str("evt", "cache.set").
		Msg("")

	expire := time.Time{}

	if ordC.ttlOrder > 0 {
		expire = time.Now().Add(ordC.ttlOrder)
	}

	// Lock - we fill the map
	ordC.mu.Lock()
	defer ordC.mu.Unlock()

	if len(ordC.m) >= ordC.cap {
		ordC.clearExpired()
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
	ordC.logger.Debug().
		Str("evt", "cache.get").
		Msg("")

	ordC.mu.RLock()
	defer ordC.mu.RUnlock()

	item, ok := ordC.m[orderUID]
	if !ok {
		return nil
	}

	b := make([]byte, len(item.info))
	copy(b, item.info)
	return b
}

func (ordC *orderCache) clearExpired() {
	ordC.logger.Debug().
		Str("evt", "cache.clearExpired").
		Msg("")

	// ttlOrder == 0 --> no ttl
	if ordC.ttlOrder == 0 {
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
	ordC.logger.Debug().
		Str("evt", "cache.fillCache").
		Msg("")

	oDB, err := ordC.db.GetN(pCtx, ordC.cap)
	if err != nil {
		return err
	}

	//OrderDataDb -> OrderData
	orders := validation.ConvertSlice(oDB)

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
		ordC.logger.Info().Msg("db.load.empty")
	} else {
		ordC.logger.Info().
			Int("Amount", i).
			Msg("db.load.ok")
	}
	return nil
}
