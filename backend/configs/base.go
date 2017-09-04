package configs

import (
	"path/filepath"
)

type AwsConfigs struct {
	ALLOW_ORIGINS       string
	BASE_MONGOURL       string
	BASE_DOMAIN         string
	REDIS_URL           string
	MONGO_AUTH_USERNAME string
	MONGO_AUTH_PASSWORD string
}

const (
	DEBUG               = true
	ALLOW_ORIGINS       = "http://localhost:8080"
	BASE_MONGOURL       = "localhost:27017"
	BASE_DOMAIN         = "localhost"
	MONGO_AUTH_USERNAME = "root"
	MONGO_AUTH_PASSWORD = "2901307001"
	AWS_REGION          = "ap-northeast-2"
	AWS_S3_BUCKET       = "xyp-s3"
	STATIC_S3_STORAGE   = false
)

var (
	AWS_CONFIGS = AwsConfigs{
		ALLOW_ORIGINS:       "http://www.xieyuanpeng.com",
		BASE_MONGOURL:       "localhost:27017",
		BASE_DOMAIN:         "xieyuanpeng.com",
		REDIS_URL:           "xyp-redis.whqvsp.0001.apn2.cache.amazonaws.com:6379",
		MONGO_AUTH_USERNAME: "xyp",
		MONGO_AUTH_PASSWORD: "pengye2901307001",
	}
	// TODO: not good here, use a configuration tool like viper later
	BACKEND_ROOT, _ = filepath.Abs("/home/xyp/go/src/github.com/pengye91/xieyuanpeng.in")
	MEDIA_ROOT      = filepath.Join(BACKEND_ROOT, "public", "media")
	IMAGE_ROOT      = filepath.Join(BACKEND_ROOT, "static", "images")
	HTML_ROOT      = filepath.Join(BACKEND_ROOT, "static", "html")
)
