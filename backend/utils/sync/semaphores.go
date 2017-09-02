package sync

import (
	"time"

	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
	"github.com/satori/go.uuid"
)

func AcquireSemaphoreBasedOnTime(semaName string, limit int, semaExpire time.Duration) (string, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()
	now := time.Now().UnixNano()

	conn.Send("MULTI")
	conn.Send("ZREMRANGEBYSCORE", semaName, 0, now-int64(semaExpire)) // clean expired semaphore
	conn.Send("ZADD", semaName, now, identifier)                      // trying to acquire a semaphore
	conn.Send("ZRANK", semaName, identifier)
	if replies, err := redis.Ints(conn.Do("EXEC")); err != nil {
		log.LoggerSugar.Errorw("sync semaphores AcquireSemaphoreBasedOnTime Error",
			"module", "application: redis",
			"error", err,
		)
		return "", err
	} else if replies[len(replies)-1] >= limit {
		conn.Do("ZREM", semaName, identifier) // delete the semaphore
		log.LoggerSugar.Warnw("sync semaphores AcquireSemaphoreBasedOnTime warning",
			"module", "application: redis",
			"warn", "trying to acquire sema but no more can be acquired now",
		)
		return "", errors.New("acquire sema failed because no more can be acquired now.")
	}
	return identifier, nil
}

func ReleaseSemaphoreBasedOnTime(semaName string, identifier string) (int, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ZREM", semaName, identifier))
	if err != nil {
		log.LoggerSugar.Errorw("sync semaphores ReleaseSemaphoreBasedOnTime ZREM Error",
			"module", "application: redis",
			"error", err,
		)
		return 0, err
	}
	return reply, nil
}

// if no expiration need, only a counter zset is totally functional.
// return identifier string and error:
// 	1. Acquire success: return identifier string and nil.
//	2. Acquire fail: return "" and the corresponding error.
func AcquireFairSemaphore(semaName string, limit int, semaExpire time.Duration) (string, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()
	now := time.Now().UnixNano()
	counterZset := semaName + ":owner"
	counter := semaName + ":counter"

	conn.Send("MULTI")
	conn.Send("ZREMRANGEBYSCORE", semaName, 0, now-int64(semaExpire))                // clean expired semaphore
	conn.Send("ZINTERSTORE", counterZset, 2, counterZset, semaName, "WEIGHTS", 1, 0) // clean expired semaphore, 妙啊
	conn.Send("INCR", counter)
	replies, err0 := redis.Ints(conn.Do("EXEC"))
	if err0 != nil {
		log.LoggerSugar.Errorw("sync semaphores AcquireFairSemaphore Error",
			"module", "application: redis",
			"error", err0,
		)
		return "", err0
	}
	ctr := replies[len(replies)-1]

	conn.Send("MULTI")
	conn.Send("ZADD", semaName, now, identifier)
	conn.Send("ZADD", counterZset, ctr, identifier)
	conn.Send("ZRANK", counterZset, identifier)

	replies1, err1 := redis.Ints(conn.Do("EXEC"))
	if err1 != nil {
		log.LoggerSugar.Errorw("sync semaphores AcquireFairSemaphore Error",
			"module", "application: redis",
			"error", err1,
		)
		return "", err1
	} else if replies1[len(replies1)-1] >= limit {
		conn.Do("ZREM", semaName, identifier) // delete the semaphore
		conn.Do("ZREM", counterZset, identifier)
		log.LoggerSugar.Warnw("sync semaphores AcquireFairSemaphore warning",
			"module", "application: redis",
			"warn", "acquire sema failed because no more can be acquired now",
		)
		return "", errors.New("acquire sema failed because no more can be acquired now.")
	}
	return identifier, nil
}

// return:
// 	- If release succeed, return true and nil.
// 	- If failed, return false and error.
func ReleaseFairSemaphore(semaName string, identifier string) (bool, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	counterZset := semaName + ":owner"

	conn.Send("ZREM", semaName, identifier)
	conn.Send("ZREM", counterZset, identifier)

	_, err := redis.Ints(conn.Do(""))
	if err != nil {
		log.LoggerSugar.Errorw("sync semaphores ReleaseFairSemaphore error",
			"module", "application: redis",
			"error", err,
		)
		return false, err
	}
	return true, nil
}

// if the identifier is expired, return false.
// else, just update the score of the identifier in semaName zset.
func RefreshFairSemaphore(semaName string, identifier string) (bool, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()
	now := time.Now().UnixNano()

	if reply, err := redis.Int(conn.Do("ZADD", semaName, now, identifier)); err != nil {
		log.LoggerSugar.Errorw("sync semaphores RefreshFairSemaphore error",
			"module", "application: redis",
			"error", err,
		)
		return false, err
	} else if reply == 1 {

		// reply == 1 means that the identifier is not in the semaName zset,
		// which indicates that the sema is timeout.
		// So return false and delete it from the semaName just added.
		ReleaseFairSemaphore(semaName, identifier)
		log.LoggerSugar.Warnw("sync semaphores RefreshFairSemaphore warning",
			"module", "application: redis",
			"warn", "The semaphore you required is timeout.",
		)
		return false, errors.New("The semaphore you required is timeout.")
	}
	return true, nil
}

func AcquireSemaphoreWithLock(semaName string, limit int, semaExpire time.Duration) (string, error) {
	conn := cache.GlobalLockRedisPool.Get()
	defer conn.Close()

	// true key name of the lock is "lock:{semaName}" instead of just sameName.
	lockId, err := AcquireLock(semaName, 10*time.Millisecond)
	if err != nil {
		log.LoggerSugar.Errorw("sync semaphore AcquireSemaphoreWithLock AcquireLock error",
			"module", "application: redis",
			"error", err,
		)
		return "", err
	}
	if released := ReleaseLock(semaName, lockId); !released {
		log.LoggerSugar.Warn("sync semaphore AcquireSemaphoreWithLock releaseLock warn",
			"module", "application: redis",
			"warn", "Release Lock Failed",
		)
		return "", errors.New("Release Lock Failed")
	}
	id, err0 := AcquireFairSemaphore(semaName, limit, semaExpire)
	if err0 != nil {
		log.LoggerSugar.Errorw("sync semaphore AcquireSemaphoreWithLock AcquireFairSemaphore error",
			"module", "application: redis",
			"error", err0,
		)
		return "", err0
	}
	return id, nil
}
