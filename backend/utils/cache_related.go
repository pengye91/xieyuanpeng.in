package utils

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
)

var (
	dialConnectTimeout       redis.DialOption = redis.DialConnectTimeout(configs.RedisTimeout)
	dialUserDatabase         redis.DialOption = redis.DialDatabase(configs.UserRelateDatabase)
	dialStatisticDatabase    redis.DialOption = redis.DialDatabase(configs.StatisticRelateDatabase)
	dialIPDatabase           redis.DialOption = redis.DialDatabase(configs.IPRelateDatabase)
	dialLockDatabase         redis.DialOption = redis.DialDatabase(configs.LockRelateDatabase)
	dialReadTimeout          redis.DialOption = redis.DialReadTimeout(configs.RedisTimeout)
	dialWriteTimeout         redis.DialOption = redis.DialWriteTimeout(configs.RedisTimeout)
	GlobalUserRedisPool      *redis.Pool      = XypRedisUserPool(configs.REDIS_URL)
	GlobalStatisticRedisPool *redis.Pool      = XypRedisStatisticPool(configs.REDIS_URL)
	GlobalIPRedisPool        *redis.Pool      = XypRedisIPPool(configs.REDIS_URL)
	GlobalLockRedisPool      *redis.Pool      = XypRedisLockPool(configs.REDIS_URL)
)

func XypRedisLockPool(addr string) *redis.Pool {
	dialOptions := []redis.DialOption{
		dialConnectTimeout, dialLockDatabase, dialReadTimeout, dialWriteTimeout,
	}
	return &redis.Pool{
		// Basic dial configuration
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, dialOptions...)
		},
		MaxIdle:     configs.RedisPoolMaxIdleNumber,
		IdleTimeout: configs.RedisPoolIdleTimeout,
	}
}

func XypRedisUserPool(addr string) *redis.Pool {
	dialOptions := []redis.DialOption{
		dialConnectTimeout, dialUserDatabase, dialReadTimeout, dialWriteTimeout,
	}
	return &redis.Pool{
		// Basic dial configuration
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, dialOptions...)
		},
		MaxIdle:     configs.RedisPoolMaxIdleNumber,
		IdleTimeout: configs.RedisPoolIdleTimeout,
	}
}

func XypRedisStatisticPool(addr string) *redis.Pool {
	dialOptions := []redis.DialOption{
		dialConnectTimeout, dialStatisticDatabase, dialReadTimeout, dialWriteTimeout,
	}
	return &redis.Pool{
		// Basic dial configuration
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, dialOptions...)
		},
		MaxIdle:     configs.RedisPoolMaxIdleNumber,
		IdleTimeout: configs.RedisPoolIdleTimeout,
	}
}

func XypRedisIPPool(addr string) *redis.Pool {
	dialOptions := []redis.DialOption{
		dialConnectTimeout, dialIPDatabase, dialReadTimeout, dialWriteTimeout,
	}
	return &redis.Pool{
		// Basic dial configuration
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, dialOptions...)
		},
		MaxIdle:     configs.RedisPoolMaxIdleNumber,
		IdleTimeout: configs.RedisPoolIdleTimeout,
	}
}
