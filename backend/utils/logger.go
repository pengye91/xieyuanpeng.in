package utils

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/kataras/iris.v6"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

type GinLoggerMiddleware struct {
	path string
	latency time.Duration
	method string
	statusCode int
	clientIP string
}

func (l *GinLoggerMiddleware) Serve(c *iris.Context) {
	// Start timer
	start := time.Now()
	l.path = c.Path()

	// Process request
	c.Next()

	// Log only when path is not being skipped
	end := time.Now()
	l.latency = end.Sub(start)

	l.clientIP = c.RemoteAddr()
	l.method = c.Method()
	statusCode := c.ResponseWriter.StatusCode()
	var statusColor, methodColor string
	if true {
		statusColor = colorForStatus(statusCode)
		methodColor = colorForMethod(l.method)
	}

	fmt.Fprintf(os.Stdout, "[xieyuanpeng.in] %v |%s %1d %s| %8v | %s |%s %s %s |%s\n",
		end.Format("01/02-15:04:05"),
		statusColor, statusCode, reset,
		l.latency,
		l.clientIP,
		methodColor, l.method, reset,
		l.path,
	)
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return green
	case code >= 300 && code < 400:
		return white
	case code >= 400 && code < 500:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

func New() iris.HandlerFunc {
	c := &GinLoggerMiddleware{}
	return c.Serve
}
