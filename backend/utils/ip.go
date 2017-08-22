package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/garyburd/redigo/redis"
)

// convert IP to redis zset score type
func ipToScore(ip string) (score int64) {
	score = 0
	ip = strings.Split(ip, "/")[0]

	for _, v := range strings.Split(ip, ".") {
		if field, err := strconv.ParseInt(v, 10, 64); err != nil {
			fmt.Println(err)
		} else {
			score = score * 256 + field
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
		fmt.Println(err)
		return err
	}
	ch := parseCSV(file)

	header := <-ch
	fmt.Println(header)
	count := 0
	conn.Do("DEL", "IP2CityID")

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
		if _, err := conn.Do("ZADD", "IP2CityID", startIPScore, cityID); err != nil {
			fmt.Println("ZADD cityID startIPScore err:")
			fmt.Println(err)
			return err
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
		fmt.Println(err)
		return err
	}
	ch := parseCSV(file)

	header := <-ch
	fmt.Println(header)
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
			fmt.Println(err)
			return err
		} else {
			conn.Do("HSET", "CityID2CityName:", cityID, cityInfoJson)
		}
	}
	return nil
}

func FindCityByIP(ip string) ([]string, error) {
	conn := GlobalIPRedisPool.Get()
	conn.Close()
	var (
		cityID string
		cityInfo []string
	)

	ipScore := ipToScore(ip)
	if reply, err := redis.Strings(conn.Do("ZREVRANGEBYSCORE", "IP2CityID", ipScore, 0, "LIMIT", 0, 1)); err != nil {
		fmt.Println(err)
		return cityInfo, err
	} else {
		cityID = strings.Split(reply[0], "_")[0]
	}

	if reply, err := redis.Bytes( conn.Do("HGET", "CityID2CityName", cityID)); err != nil {
		fmt.Println(err)
		return cityInfo, err
	} else {
		err := json.Unmarshal(reply, cityInfo)
		if err != nil {
			fmt.Println(err)
			return cityInfo, err
		}
		return cityInfo, nil
	}
}

