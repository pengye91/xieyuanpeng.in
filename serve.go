package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"gopkg.in/kataras/iris.v5"
)

func main() {
	iris.Use(logger.New())
	iris.Use(cors.Default())
	iris.Post("/v1/auth/test1", func (ctx *iris.Context) {
		ctx.JSON(200, "xixi")
	})
	iris.Get("/v1/auth/test", func (ctx *iris.Context) {
		ctx.JSON(200, "xixi")
	})
	iris.Listen(":8001")
}


