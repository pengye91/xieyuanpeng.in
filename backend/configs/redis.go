package configs

import "time"

const (
	UserRelateDatabase      = 2
	StatisticRelateDatabase = 3
	IPRelateDatabase        = 4
	LockRelateDatabase      = 5
	MQRelateDatabase        = 6
	RedisTimeout            = 10 * time.Second
	RedisPoolMaxIdleNumber  = 20
	// 0 means never close idle connections
	RedisPoolIdleTimeout = 0
	REDIS_URL            = "192.168.2.112:6379"
	//REDIS_URL            = "192.168.0.106:6379"
)
