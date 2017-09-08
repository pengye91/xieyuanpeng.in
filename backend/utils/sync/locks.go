package sync

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
	"github.com/satori/go.uuid"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

func AcquireLock(lockname string, acquireTimeout time.Duration) (string, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()

	endTime := time.Now().UnixNano() + int64(acquireTimeout)

	for time.Now().UnixNano() < endTime {
		if reply, err := redis.Bool(conn.Do("SET", "lock:"+lockname, identifier, "NX")); err != nil {
			log.LoggerSugar.Errorw("redis lock acquireLock SET Error",
				"module", "redis",
				"error", err,
			)
			time.Sleep(1 * time.Millisecond)
		} else {
			log.LoggerSugar.Infow("redis lock acquireLock SET Info",
				"module", "redis",
				"reply", reply,
			)
			return identifier, nil
		}
	}
	return "", errors.New("TimeOut")
}

func AcquireLockWithTimeout(lockname string, acquireTimeout time.Duration, lockTimeout time.Duration) (string, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()

	endTime := time.Now().UnixNano() + int64(acquireTimeout)

	for time.Now().UnixNano() < endTime {
		if reply, err := redis.Bool(conn.Do("SET", "lock:"+lockname, identifier, "EX", lockTimeout, "NX")); err != nil {
			log.LoggerSugar.Errorw("sync lock AcquireLockWithTimeout SET lock error",
				"module", "redis",
				"error", err,
			)
			time.Sleep(1 * time.Millisecond)
		} else {
			log.LoggerSugar.Infow("redis lock acquireLock SET Info",
				"module", "redis",
				"reply", reply,
			)
			return identifier, nil
		}
	}
	return identifier, errors.New("TimeOut")
}

func ReleaseLock(lockname string, identifier string) bool {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	lockname = "lock:" + lockname

	for true {
		conn.Send("WATCH", lockname)
		if reply, _ := redis.String(conn.Do("GET", lockname)); reply == identifier {
			conn.Send("MULTI")
			conn.Send("DEL", lockname)
			if reply, err := conn.Do("EXEC"); err != nil {

				// err != nil means that the transaction failed.
				// should repeat the operation
				log.LoggerSugar.Errorw("sync lock ReleaseLock Error",
					"module", "application: redis",
					"error", err,
				)
				continue
			} else {
				log.LoggerSugar.Infow("redis lock acquireLock SET Info",
					"module", "redis",
					"reply", reply,
				)
				return true
			}
		}
		break
	}
	return false
}
