package utils

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/kataras/go-errors"
	"github.com/satori/go.uuid"
)

func AcquireSemaphoreBasedOnTime(semaName string, limit int, semaExpire int64) (string, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()
	now := time.Now().UnixNano()

	conn.Send("MULTI")
	conn.Send("ZREMRANGEBYSCORE", semaName, 0, now-(semaExpire*1e9)) // clean expired semaphore
	conn.Send("ZADD", semaName, now, identifier)                     // trying to acquire a semaphore
	conn.Send("ZRANK", semaName, identifier)
	if replies, err := redis.Ints(conn.Do("EXEC")); err != nil {
		fmt.Println("EXEC err:")
		fmt.Println(err)
		return "", err
	} else if replies[len(replies)-1] >= limit {
		conn.Do("ZREM", semaName, identifier) // delete the semaphore
		return "", errors.New("acquire sema failed because no more can be acquired now.")
	} else {
		fmt.Println(replies)
	}

	return identifier, nil
}

func ReleaseSemaphoreBasedOnTime(semaName string, identifier string) (int, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ZREM", semaName, identifier))
	if err != nil {
		fmt.Println("EXEC err:")
		fmt.Println(err)
		return 0, err
	}

	fmt.Println(reply)
	return reply, nil
}

// if no expiration need, only a counter zset is totally functional.
// return identifier string and error:
// 	1. Acquire success: return identifier string and nil.
//	2. Acquire fail: return "" and the corresponding error.
func AcquireFairSemaphore(semaName string, limit int, semaExpire int64) (string, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()
	now := time.Now().UnixNano()
	counterZset := semaName + ":owner"
	counter := semaName + ":counter"

	conn.Send("MULTI")
	conn.Send("ZREMRANGEBYSCORE", semaName, 0, now-(semaExpire*1e9))                 // clean expired semaphore
	conn.Send("ZINTERSTORE", counterZset, 2, counterZset, semaName, "WEIGHTS", 1, 0) // clean expired semaphore, 妙啊
	conn.Send("INCR", counter)
	replies, err0 := redis.Ints(conn.Do("EXEC"))
	if err0 != nil {
		fmt.Println("EXEC err:")
		fmt.Println(err0)
		return "", err0
	}
	ctr := replies[len(replies)-1]

	conn.Send("MULTI")
	conn.Send("ZADD", semaName, now, identifier)
	conn.Send("ZADD", counterZset, ctr, identifier)
	conn.Send("ZRANK", counterZset, identifier)

	replies1, err1 := redis.Ints(conn.Do("EXEC"))
	if err1 != nil {
		fmt.Println("EXEC err:")
		fmt.Println(err1)
		return "", err1
	} else if replies1[len(replies1)-1] >= limit {
		conn.Do("ZREM", semaName, identifier) // delete the semaphore
		conn.Do("ZREM", counterZset, identifier)
		return "", errors.New("acquire sema failed because no more can be acquired now.")
	}
	return identifier, nil
}

// return:
// 	1. If release succeed, return true and nil.
// 	2. If failed, return false and error.
func ReleaseFairSemaphore(semaName string, identifier string) (bool, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	counterZset := semaName + ":owner"

	conn.Send("ZREM", semaName, identifier)
	conn.Send("ZREM", counterZset, identifier)

	reply, err := redis.Ints(conn.Do(""))
	if err != nil {
		fmt.Println("EXEC err:")
		fmt.Println(err)
		return false, err
	}

	fmt.Println(reply)
	return true, nil
}

// if the identifier is expired, return false.
// else, just update the score of the identifier in semaName zset.
func RefreshFairSemaphore(semaName string, identifier string) (bool, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()
	now := time.Now().UnixNano()

	if reply, err := redis.Int(conn.Do("ZADD", semaName, now, identifier)); err != nil {
		fmt.Println("Do ZADD err:")
		fmt.Println(err)
		return false, err
	} else if reply == 1 {

		// reply == 1 means that the identifier is not in the semaName zset,
		// which indicates that the sema is timeout.
		// So return false and delete it from the semaName just added.
		ReleaseFairSemaphore(semaName, identifier)
		return false, errors.New("The semaphore you required is timeout.")
	}
	return true, nil
}

func AcquireSemaphoreWithLock(semaName string, limit int, semaExpire int64) (string, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	// true key name of the lock is "lock:{semaName}" instead of just sameName.
	_, err := AcquireLock(semaName, 0.01)
	if err != nil {
		fmt.Println("AcquireLock err:")
		fmt.Println(err)
		return "", err
	}
	id, err0 := AcquireFairSemaphore(semaName, limit, semaExpire)
	if err0 != nil {
		fmt.Println("AcquireFairSemaphore err:")
		fmt.Println(err0)
		return "", err0
	}
	return id, nil
}
