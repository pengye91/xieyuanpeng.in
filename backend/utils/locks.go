package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
)

func AcquireLock(lockname string, acquireTimeout int64) (string, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()

	endTime := time.Now().Unix() + acquireTimeout

	for time.Now().Unix() < endTime {
		if reply, err := redis.Bool(conn.Do("SET", "lock:"+lockname, identifier, "NX")); err != nil {
			fmt.Println("setnx wrong:")
			fmt.Println(err)
			time.Sleep(1 * time.Millisecond)
		} else {
			fmt.Println(reply)
			return identifier, nil
		}
	}
	return "", errors.New("TimeOut")
}

func AcquireLockWithTimeout(lockname string, acquireTimeout int64, lockTimeout int64) (string, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := uuid.NewV4().String()

	endTime := time.Now().Unix() + acquireTimeout

	for time.Now().Unix() < endTime {
		if reply, err := redis.Bool(conn.Do("SET", "lock:"+lockname, identifier, "EX", lockTimeout, "NX")); err != nil {
			fmt.Println("setnx wrong:")
			fmt.Println(err)
			time.Sleep(1 * time.Millisecond)
		} else {
			fmt.Println(reply)
			return identifier, nil
		}
	}
	return identifier, errors.New("TimeOut")
}

func ReleaseLock(lockname string, identifier string) bool {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	lockname = "lock:" + lockname

	for true {
		conn.Send("WATCH", lockname)
		if reply, _ := redis.String(conn.Do("GET", lockname)); reply == identifier {
			conn.Send("MULTI")
			conn.Send("DEL", lockname)
			if reply, err := conn.Do("EXEC"); err != nil {
				fmt.Println("EXEC err:")
				fmt.Println(err)
				continue
			} else {
				fmt.Println(reply)
				return true
			}
		}
		break
	}
	return false
}
