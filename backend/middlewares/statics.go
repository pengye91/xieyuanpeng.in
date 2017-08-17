package middlewares

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
)

func TotalHitMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn := utils.GlobalStaticRedisPool.Get()
		defer conn.Close()

		if reply, err := redis.Int(conn.Do("INCR", "TotalHit")); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(reply)
		}
		ctx.Next()
	}
}
