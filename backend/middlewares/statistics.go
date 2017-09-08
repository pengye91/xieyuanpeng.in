package middlewares

import (
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

var (
	timeSliceCountChan = make(chan map[string]int64, 10)
)

// GlobalStatisticsMiddleware put all kinds of global statistics data into Redis
func GlobalStatisticsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		go TotalHitsCount()
		go TimeSliceCount()
		//go GetTimeSliceCount(&timeSliceCountChan)
		ctx.Next()
	}
}

func TotalHitsCount() {
	conn := cache.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	if _, err := redis.Int(conn.Do("INCR", "TotalHit")); err != nil {
		log.LoggerSugar.Errorw("statistics TotalHitsCount INCR Error",
			"module", "redis",
			"error", err,
		)
	}
}

// update the hit number when got hit
// Counter.
func TimeSliceCount() {
	conn := cache.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	now := time.Now().Unix()
	for _, pre := range configs.PRECISION {
		// why starting with 08:00?
		preNowInUnix := pre * (now / pre)
		preNowHuman := time.Unix(preNowInUnix, 0).Format("2006/01/02/15-04-05")
		hash := strconv.FormatInt(pre, 10) + ":hits"
		conn.Send("ZADD", "known:", 0, hash)
		conn.Send("HINCRBY", "count:"+hash, preNowHuman, 1)
	}

	// use Send and Do to implement pipeline
	if _, err := conn.Do(""); err != nil {
		log.LoggerSugar.Errorw("statistic Do TimeSliceCount Error",
			"module", "redis",
			"error", err,
		)
		return
	}
}

// Get allTimeSliceCount from Redis, not sorted Counter.
func GetTimeSliceCount(c *chan map[string]int64) {
	conn := cache.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	for _, pre := range configs.PRECISION {
		// why starting with 08:00?

		// "count" is first type layer.
		// "pre" is second layer.
		// "hits" is type.
		hash := strconv.FormatInt(pre, 10) + ":hits"

		// Use Int64Map to get "HGETALL" results
		if timeSliceCountMaps, err := redis.Int64Map(conn.Do("HGETALL", "count:"+hash)); err != nil {
			log.LoggerSugar.Errorw("statistics GetTimeSliceCount HGETALL Error",
				"module", "redis",
				"error", err,
			)
			return
		} else {
			*c <- timeSliceCountMaps
		}
	}
	go ConsumeFromChan(c)
}

// Counter helper.
func ConsumeFromChan(c *chan map[string]int64) {
	// TODO: replace this into real logic
	for t := range *c {
		log.LoggerSugar.Infow("statistics ConsumeFromChan Info",
			"module", "redis and application",
			"info", t,
		)
		//for timeSlice, hits := range t {
		//	fmt.Printf("HITS IN %s: %d \n", timeSlice, hits)
		//}
	}
}
