package utils

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/satori/go.uuid"
)

func AcquireLock(lockname string, timeout int64) (string, error) {
	conn := GlobalLockRedisPool.Get()
	defer conn.Close()

	identifier := string(uuid.NewV4())

	endTime := time.Now().Unix() + timeout

	for time.Now().Unix() < endTime {
		if reply, err := redis.Bool(conn.Do("SETNX", "lock:"+lockname, identifier)); err != nil {
			fmt.Println("setnx wrong:")
			fmt.Println(err)
			return identifier, err
		} else {
			fmt.Println(reply)
			return identifier, nil
		}
	}
}
