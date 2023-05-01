package rediscache

import "time"

type redisConfig struct {
	URL         []string
	MaxRetries  int
	DialTimeout time.Duration
	IdleTimeout time.Duration
}
