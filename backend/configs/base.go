package configs

type AwsConfigs struct {
	ALLOW_ORIGINS       string
	BASE_MONGOURL       string
	BASE_DOMAIN         string
	REDIS_URL           string
	MONGO_AUTH_USERNAME string
	MONGO_AUTH_PASSWORD string
}

const (
	ALLOW_ORIGINS       = "http://localhost:8080"
	BASE_MONGOURL       = "localhost:27017"
	BASE_DOMAIN         = "localhost"
	REDIS_URL           = "localhost:6379"
	MONGO_AUTH_USERNAME = "xyp"
	MONGO_AUTH_PASSWORD = "xxyypp"
)

var (
	AWS_CONFIGS = AwsConfigs{
		ALLOW_ORIGINS:       "http://www.xieyuanpeng.com",
		BASE_MONGOURL:       "localhost:27017",
		BASE_DOMAIN:         "xieyuanpeng.com",
		REDIS_URL:           "xyp-redis.whqvsp.0001.apn2.cache.amazonaws.com:6379",
		MONGO_AUTH_USERNAME: "xyp",
		MONGO_AUTH_PASSWORD: "2901307001",
	}
)
