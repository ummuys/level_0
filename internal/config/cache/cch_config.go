package config

import "time"

type CacheEnv struct {
	Capacity int
	TTLOrder time.Duration
}
