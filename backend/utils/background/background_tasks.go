package background

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

// Create a background goroutine deleting Redis keys periodically.
func CleanTimeSlice() {
	conn := cache.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	for {
		start := time.Now()
		hashes, _ := redis.Strings(conn.Do("ZRANGE", "known:", 0, -1))
		passes := 0
		for _, hash := range hashes {
			if hash == "" {
				break
			}
			hashCounter := "count:" + hash
			prec, _ := strconv.Atoi(strings.Split(hash, ":")[0])

			preMinute := 1
			if (prec / 60) != 0 {
				preMinute = prec / 60
			}

			if (passes % preMinute) != 0 {
				continue
			}

			timeSlices, err := redis.Strings(conn.Do("HKEYS", hashCounter))
			if err != nil {
				log.LoggerSugar.Errorw("utils background tasks HKEYS error",
					"module", "redis",
					"error", err,
				)
			} else {
				sort.Strings(timeSlices)
				log.LoggerSugar.Infow("utils background tasks HKEYS Info",
					"module", "application and redis",
					"timeSlices", timeSlices,
				)
			}

			if len(timeSlices) > configs.SAMPLE_COUNT {
				log.LoggerSugar.Infow("utils background tasks HKEYS Info",
					"module", "application and redis",
					"info", "greater than sample count",
					"hashCounter", hashCounter,
					"HDEL", timeSlices[:len(timeSlices)-configs.SAMPLE_COUNT],
				)
				if reply, err := redis.Int(conn.Do("HDEL", redis.Args{}.Add(hashCounter).AddFlat(timeSlices[:len(timeSlices)-configs.SAMPLE_COUNT])...)); err != nil {
					log.LoggerSugar.Errorw("utils background tasks HDEL error",
						"module", "redis",
						"error", err,
					)
				} else {
					log.LoggerSugar.Infow("utils background tasks HDEL Info",
						"module", "application and redis",
						"reply", reply,
					)
				}
			} else {
				log.LoggerSugar.Infow("utils background tasks HDEL Info",
					"module", "application and redis",
					"info", "less than sample count, no need to clean",
				)
			}
		}
		passes += 1
		duration := math.Min(time.Since(start).Seconds(), 60)
		time.Sleep(time.Duration(math.Max(60-duration, 1)) * time.Second)
	}
}

func UpdateContextStats(context string, statsType string, value string) {

}
