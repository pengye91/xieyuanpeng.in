package utils

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
)

var (
	dialConnectTimeout    redis.DialOption = redis.DialConnectTimeout(configs.RedisTimeout)
	dialUserDatabase      redis.DialOption = redis.DialDatabase(configs.UserRelateDatabase)
	dialStaticDatabase    redis.DialOption = redis.DialDatabase(configs.StaticRelateDatabase)
	dialReadTimeout       redis.DialOption = redis.DialReadTimeout(configs.RedisTimeout)
	dialWriteTimeout      redis.DialOption = redis.DialWriteTimeout(configs.RedisTimeout)
	GlobalUserRedisPool   *redis.Pool      = XypRedisUserPool(configs.REDIS_URL)
	GlobalStaticRedisPool *redis.Pool      = XypRedisStaticPool(configs.REDIS_URL)
)

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

func XypRedisStaticPool(addr string) *redis.Pool {
	dialOptions := []redis.DialOption{
		dialConnectTimeout, dialStaticDatabase, dialReadTimeout, dialWriteTimeout,
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
