package main

import (
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
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
				Debug:          true,
			},
		),
	)
<<<<<<< HEAD
	// logger
	app.Use(utils.New())

	app.StaticServe("../xiyuanpeng_front/public", "/static")
	v1 := app.Party("/v1")
	{
		b := v1.Party("/auth1")
		{
			b.Post("/register", func(c *iris.Context) {
				c.JSON(200, "v1.auth.register.post")
			})
			b.Post("/login", func(c *iris.Context) {
				c.JSON(200, "v1.auth.login.post")
			})
			b.Get("/login", func(c *iris.Context) {
				c.JSON(200, "v1.auth.login.get")
			})
			b.Get("/register", func(c *iris.Context) {
				retUrl := "xixi"
				c.JSON(200, retUrl)
			})
		}
		authRoute(v1)
	}
	app.Listen(":8001")
}

func authRoute(app *iris.Router) {
	a := app.Party("/auth")
	{
		a.Post("/register", func(c *iris.Context) {
			c.JSON(200, "v1.auth.register.post")
		})
		a.Post("/login", func(c *iris.Context) {
			c.JSON(200, "v1.auth.login.post")
		})
		a.Get("/login", func(c *iris.Context) {
			c.JSON(200, "v1.auth.login.get")
		})
		a.Get("/register", func(c *iris.Context) {
			c.JSON(200, "v1.auth.register.get")
		})
	}
=======
>>>>>>> dev
}
