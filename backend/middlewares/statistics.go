package middlewares

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
)

var (
	timeSliceCountChan = make(chan map[string]int64, 10)
)

// GlobalStatisticsMiddleware put all kinds of global statistics data into Redis
func GlobalStatisticsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		go TotalHitsCount()
		go TimeSliceCount()
		go GetTimeSliceCount(&timeSliceCountChan)
		//go ConsumeFromChan(&timeSliceCountChan)
		ctx.Next()
	}
}

func TotalHitsCount() {
	conn := utils.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	if _, err := redis.Int(conn.Do("INCR", "TotalHit")); err != nil {
		fmt.Println(err)
	}
}

// update the hit number when got hit
func TimeSliceCount() {
	conn := utils.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	now := time.Now().Unix()
	for _, pre := range configs.PRECISION {
		// why starting with 08:00?
		preNowInUnix := pre * (now / pre)
		preNowHuman := time.Unix(preNowInUnix, 0).Format("2006/01/02/15-04-05")
		hash := "from-" + preNowHuman + ":hits"
		conn.Send("ZADD", "known:", 0, hash)
		conn.Send("HINCRBY", "count:"+hash, preNowHuman, 1)
	}

	// use Send and Do to implement pipeline
	if _, err := conn.Do(""); err != nil {
		fmt.Println(err)
	}
}

// Get allTimeSliceCount from Redis, not sorted
func GetTimeSliceCount(c *chan map[string]int64) {
	conn := utils.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	now := time.Now().Unix()
	for i, pre := range configs.PRECISION {
		fmt.Println(i)
		// why starting with 08:00?
		preNowInUnix := pre * (now / pre)
		preNowHuman := time.Unix(preNowInUnix, 0).Format("2006/01/02/15-04-05")
		hash := "from-" + preNowHuman + ":hits"

		// Use Int64Map to get "HGETALL" results
		if timeSliceCountMaps, err := redis.Int64Map(conn.Do("HGETALL", "count:"+hash)); err != nil {
			fmt.Println(err)
		} else {
			*c <- timeSliceCountMaps
		}
	}

	go ConsumeFromChan(c)
}

func ConsumeFromChan(c *chan map[string]int64) {
	// TODO: replace this into real logic
	for t := range *c {
		fmt.Println(t)
		for timeSlice, hits := range t {
			fmt.Printf("HITS IN %s: %d \n", timeSlice, hits)
		}
	}
}
