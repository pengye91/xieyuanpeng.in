package mq

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
)

// Use RPUSH to push an email content into a redis list.
// Return true if push succeed, otherwise false.
func SendEmailViaQueue(emailContent map[string]string) (bool, error) {
	conn := cache.GlobalMQRedisPool.Get()
	defer conn.Close()

	emailContentJSON, err := json.Marshal(emailContent)
	if err != nil {
		fmt.Println("JSON Marshal error:")
		fmt.Println(err)
		return false, err
	}
	reply, RpushErr := conn.Do("RPUSH", "QUEUE:EMAIL", emailContentJSON)
	if RpushErr != nil {
		fmt.Println("RPUSH error:")
		fmt.Println(err)
		return false, RpushErr
	}
	fmt.Println(reply)
	return true, nil
}

// This is tricky, the GlobalMQPool.Get() method will return a connection which will get closed in the DialReadTimeout.
// Use a new connection returned from Dial() which sets the timeout with 0.
// This connection will never get closed until the program dies.
func ProcessEmailQueue() {
	conn, connErr := redis.Dial("tcp", configs.REDIS_URL, redis.DialDatabase(configs.MQRelateDatabase), redis.DialReadTimeout(0), redis.DialWriteTimeout(0))
	if connErr != nil {
		fmt.Println(connErr)
	}
	defer conn.Close()

	for true {
		reply, BLPOPErr := redis.ByteSlices(conn.Do("BLPOP", "QUEUE:EMAIL", 30))
		if BLPOPErr != nil {
			fmt.Println("BLPOP error:")
			fmt.Println(BLPOPErr)
			continue
		}
		fmt.Println(reply)
		if len(reply) == 0 {
			continue
		}
		var emailContent = make(map[string]string)

		JSONDecodeErr := json.Unmarshal(reply[1], &emailContent)
		if JSONDecodeErr != nil {
			fmt.Println("JSONDecoder error:")
			fmt.Println(JSONDecodeErr)
			return
		}
		fmt.Println(emailContent)
	}
}

func DelayExecute(queue string, functionName string, args map[string]interface{}, delay time.Duration) {
}

type tt struct {
	a int
	b int
}

type TT struct {
	*tt
}

func CreateTT(a int, b int) *TT {
	t := tt{a, b}
	T := TT{&t}
	fmt.Println(T)
	return &T
}

func (t TT) TestTT() {
	fmt.Printf("t.a: %v: %T\n", t.a, t.a)
	fmt.Printf("t: %v: %T\n", t, t)
	fmt.Printf("t.tt: %v: %T\n", t.tt, t.tt)
}
