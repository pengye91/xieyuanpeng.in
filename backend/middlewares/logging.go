package middlewares

import (
	//"bytes"
	//"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
	"io/ioutil"
	"bytes"
)

func GlobalLoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ip, path, method, data := requestInfo(ctx)
		ctx.Next()
		elapsed, status := responseInfo(ctx, start)
		globalLogger, _ := log.GlobalLogger(status, elapsed, ip, path, method, data)
		if status < 400 {
			globalLogger.Info("[global access]")
		} else {
			globalLogger.Error("[global error]")
		}
	}
}

func requestInfo(ctx *gin.Context) (string, string, string, []byte) {
	ip := ctx.ClientIP()
	path := ctx.Request.URL.Path
	if ctx.Request.URL.RawQuery != "" {
		path += "?"
		path += ctx.Request.URL.RawQuery
	}
	method := ctx.Request.Method
	var (
		requestBody []byte
		err         error
	)
	if ctx.Request.Body != nil {
		requestBody, err = ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			log.LoggerSugar.Errorw("getRequestInfo read request body error",
				"module", "application: logging",
				"error", err,
			)
		}
		// store back to the Body
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	}
	return ip, path, method, requestBody
}

func responseInfo(ctx *gin.Context, start time.Time) (string, int) {
	elapsed := time.Since(start)
	status := ctx.Writer.Status()
	return elapsed.String(), status
}
