package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseCacheEnv() (*CacheEnv, error) {

	var sErr []string

	add := func(env string) {
		sErr = append(sErr, fmt.Sprintf("invalid env for %s", env))
	}

	capStr := os.Getenv("CACHE_CAPACITY")
	cap, err := strconv.Atoi(capStr)
	if err != nil {
		add("cache")
	}

	ttlOrderSecStr := os.Getenv("TTL_ORDER")
	ttlOrderSec, err := strconv.Atoi(ttlOrderSecStr)
	if err != nil {
		add("ttl_order")
	}

	if len(sErr) > 0 {
		return nil, fmt.Errorf(strings.Join(sErr, ", "))
	}

	return &CacheEnv{
		Capacity: cap,
		TTLOrder: time.Duration(ttlOrderSec) * time.Second,
	}, nil
}
