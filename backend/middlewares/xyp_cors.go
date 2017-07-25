package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
)

var (
	CORSMiddleware = cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{configs.ALLOWORIGIN},
		AllowAllOrigins:  false,
	})
)
