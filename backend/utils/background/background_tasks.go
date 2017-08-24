package background

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/cache"
)

// Create a background goroutine deleting Redis keys periodically.
// TODO: add logging
func CleanTimeSlice() {
	conn := cache.GlobalStatisticRedisPool.Get()
	// Never forget to close the connection
	defer conn.Close()

	for {
		start := time.Now()
		fmt.Printf("%s: %s\n", "start", start)
		hashes, _ := redis.Strings(conn.Do("ZRANGE", "known:", 0, -1))
		fmt.Printf("%s: %s\n", "hashes", hashes)
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
				fmt.Println(err)
			} else {
				sort.Strings(timeSlices)
				fmt.Printf("%s: %s\n", "timeSlices", timeSlices)
			}

			if len(timeSlices) > configs.SAMPLE_COUNT {
				fmt.Println("gt")
				fmt.Println(hashCounter)
				fmt.Printf("%s: %s\n", "addFlat", timeSlices[:len(timeSlices)-configs.SAMPLE_COUNT])
				if reply, err := redis.Int(conn.Do("HDEL", redis.Args{}.Add(hashCounter).AddFlat(timeSlices[:len(timeSlices)-configs.SAMPLE_COUNT])...)); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("%s: %s\n", "reply", reply)
				}
			} else {
				fmt.Println("lt")
			}
		}
		passes += 1
		duration := math.Min(time.Since(start).Seconds(), 60)
		time.Sleep(time.Duration(math.Max(60-duration, 1)) * time.Second)
	}
}

func UpdateContextStats(context string, statsType string, value string) {

}
