package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(
		iris.DevLogger(),
		httprouter.New(),
		cors.New(
			cors.Options{
				AllowedOrigins: []string{"*"},
				AllowedHeaders: []string{"*"},
			},
		),
	)
}
