package middlewares

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
)

var (
	precision          = []int64{3600, 86400}
	timeSliceCountChan = make(chan map[string]int64, 10)
)

// GlobalStatisticsMiddleware put all kinds of global statistics data into Redis
func GlobalStatisticsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		go TotalHitsCount()
		go TimeSliceCount()
		//GetTimeSliceCount()
		//ConsumeFromChan(&timeSliceCountChan)
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

func TimeSliceCount() {
	conn := utils.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	now := time.Now().Unix()
	for _, pre := range precision {
		// why starting with 08:00?
		preNowInUnix := pre * (now / pre)
		preNowHuman := time.Unix(preNowInUnix, 0).Format("2006-01-02-15-04-05")
		hash := "from-" + preNowHuman + ":hits"
		conn.Send("ZADD", "known:", 0, hash)
		conn.Send("HINCRBY", "count:"+hash, preNowHuman, 1)

		// Use Int64Map to get "HGETALL" results
		//if timeSliceCountMaps, err := redis.Int64Map(conn.Do("HGETALL", "count:"+hash)); err != nil {
		//	fmt.Println(err)
		//} else {
		//	 TODO: replace this into real logic
		//for timeSlice, hits := range timeSliceCountMaps {
		//	fmt.Printf("HITS IN %s: %d \n", timeSlice, hits)
		//}
		//timeSliceCountChan <- timeSliceCountMaps
		//}
	}

	if _, err := conn.Do(""); err != nil {
		fmt.Println(err)
	}
}

// Get allTimeSliceCount from Redis, not sorted
func GetTimeSliceCount() {
	fmt.Println("test1")
	conn := utils.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	now := time.Now().Unix()
	for i, pre := range precision {
		fmt.Println(i)
		// why starting with 08:00?
		preNowInUnix := pre * (now / pre)
		preNowHuman := time.Unix(preNowInUnix, 0).Format("2006/01/02/15-04-05")
		hash := "from-" + preNowHuman + ":hits"

		// Use Int64Map to get "HGETALL" results
		if timeSliceCountMaps, err := redis.Int64Map(conn.Do("HGETALL", "count:"+hash)); err != nil {
			fmt.Println(err)
		} else {
			// TODO: replace this into real logic
			for timeSlice, hits := range timeSliceCountMaps {
				fmt.Printf("HITS IN %s: %d \n", timeSlice, hits)
			}
			fmt.Println(i)
			timeSliceCountChan <- timeSliceCountMaps
		}
	}
}

func ConsumeFromChan(c *(chan map[string]int64)) {
	fmt.Println("test")
	for t := range *c {
		fmt.Println(t)
		for timeSlice, hits := range t {
			fmt.Printf("HITS IN %s: %d \n", timeSlice, hits)
		}
	}
}
