package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

func GlobalLoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ip, path, method := requestInfo(ctx)
		ctx.Next()
		elapsed, status := responseInfo(ctx, start)
		globalLogger, _ := log.GlobalLogger(status, elapsed, ip, path, method)
		if status < 400 {
			globalLogger.Info("[global access]")
		} else {
			globalLogger.Error("[global error]")
		}
	}
}

func requestInfo(ctx *gin.Context) (string, string, string) {
	ip := ctx.ClientIP()
	path := ctx.Request.URL.Path
	if ctx.Request.URL.RawQuery != "" {
		path += "?"
		path += ctx.Request.URL.RawQuery
	}
	method := ctx.Request.Method
	return ip, path, method
}

func responseInfo(ctx *gin.Context, start time.Time) (string, int) {
	elapsed := time.Since(start)
	status := ctx.Writer.Status()
	return elapsed.String(), status
}
