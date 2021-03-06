package cache

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

// convert IP to redis zset score type
func ipToScore(ip string) (score int64) {
	score = 0
	ip = strings.Split(ip, "/")[0]

	for _, v := range strings.Split(ip, ".") {
		if field, err := strconv.ParseInt(v, 10, 64); err != nil {
			log.LoggerSugar.Errorw("cache IpToScore ParseInt Error",
				"module", "application: ipToScore",
				"error", err,
			)
		} else {
			score = score*256 + field
		}
	}
	return
}

func parseCSV(c io.Reader) chan []string {
	ch := make(chan []string, 10)

	go func() {
		csvReader := csv.NewReader(c)
		defer close(ch)
		if header, err := csvReader.Read(); err != nil {
			panic(err)
		} else {
			ch <- header
		}
		for {
			rec, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.LoggerSugar.Errorw("cache ipToScore parseCSV Error",
						"module", "application: csv",
						"error", err,
					)
				}
			}

			ch <- rec
		}
	}()
	return ch
}

// The filename parameter is GeoLiteCity-Blocks.csv
func ImportIPToRedis(filename string) error {
	conn := GlobalIPRedisPool.Get()
	defer conn.Close()

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.LoggerSugar.Errorw("cache ip ImportToRedis Open csv file Error",
			"module", "application: redis",
			"error", err,
		)
		return err
	}
	ch := parseCSV(file)

	header := <-ch
	log.LoggerSugar.Infow("cache ip ImportToRedis Info: ip.csv header",
		"header", header,
	)
	count := 0
	conn.Do("DEL", "IP2CityID")
	zaddPipe := 0

	for rec := range ch {
		count++
		var (
			startIP      string
			startIPScore int64
			cityID       string
		)
		if len(rec) == 0 {
			startIP = ""
		} else {
			startIP = rec[0]
		}

		if strings.Contains(startIP, ".") {
			startIPScore = ipToScore(startIP)
		} else if v, err := strconv.ParseInt(startIP, 10, 64); err == nil {
			startIPScore = v
		}

		cityID = rec[1] + "_" + strconv.Itoa(count)

		if err := conn.Send("ZADD", "IP2CityID", startIPScore, cityID); err != nil {
			log.LoggerSugar.Errorw("cache ip importIpToRedis Error",
				"module", "redis",
				"error", err,
			)
			return err
		} else {
			zaddPipe++
		}
		if zaddPipe == 100 {
			if _, err := conn.Do(""); err != nil {
				fmt.Println("ZADD cityID startIPScore err:")
				fmt.Println(err)
				return err
			}
			zaddPipe = 0
		}
	}
	return nil
}

// The filename parameter is GeoLiteCity-Locations.csv
func ImportCitiesToRedis(filename string) error {
	conn := GlobalIPRedisPool.Get()
	defer conn.Close()

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.LoggerSugar.Errorw("cache ip ImportToRedis Open Cities csv file Error",
			"module", "application: redis",
			"error", err,
		)
		return err
	}
	ch := parseCSV(file)

	header := <-ch
	log.LoggerSugar.Infow("cache ip ImportToRedis Open Cities csv file Info",
		"module", "application: redis",
		"csv file header", header,
	)
	conn.Do("DEL", "CityID2CityName:")

	for rec := range ch {
		if len(rec) == 0 {
			continue
		}
		cityID := rec[0]
		countryName := rec[5]
		cityName := rec[len(rec)-3]
		cityInfo := []string{}
		cityInfo = append(cityInfo, cityID, countryName, cityName)

		if cityInfoJson, err := json.Marshal(cityInfo); err != nil {
			log.LoggerSugar.Errorw("cache ip importIpToRedis json.Marshal() Error",
				"module", "application: redis",
				"error", err,
			)
			return err
		} else {
			conn.Do("HSET", "CityID2CityName:", cityID, cityInfoJson)
		}
	}
	return nil
}

func FindCityByIP(ip string) ([]string, error) {
	conn := GlobalIPRedisPool.Get()
	defer conn.Close()
	var (
		cityID   string
		cityInfo []string
	)

	ipScore := ipToScore(ip)
	if reply, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", "IP2CityID", ipScore, 0, "LIMIT", 0, 1)); err != nil {
		log.LoggerSugar.Errorw("cache ip FindCityByIp ZREVRANGEBYSCORE Error",
			"module", "redis",
			"error", err,
		)
		return cityInfo, err
	} else {
		cityID = strings.Split(reply[0], "_")[0]
	}

	if reply, err := redis.Bytes(conn.Do("HGET", "CityID2CityName:", cityID)); err != nil {
		log.LoggerSugar.Errorw("cache ip FindCityByIp HGET Error",
			"module", "redis",
			"error", err,
		)
		return cityInfo, err
	} else {
		err := json.Unmarshal(reply, &cityInfo)
		if err != nil {
			log.LoggerSugar.Errorw("cache ip FindCityByIp json.Unmarshal Error",
				"module", "application: json",
				"error", err,
			)
			return cityInfo, err
		}
		return cityInfo, nil
	}
}
