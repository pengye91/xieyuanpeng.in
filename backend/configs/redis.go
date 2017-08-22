package configs

import "time"

const (
	UserRelateDatabase      = 2
	StatisticRelateDatabase = 3
	IPRelateDatabase        = 4
	RedisTimeout            = 10 * time.Second
	RedisPoolMaxIdleNumber  = 20
	// 0 means never close idle connections
	RedisPoolIdleTimeout = 0
	REDIS_URL            = "localhost:6379"
)
