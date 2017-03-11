package main

import (
	"github.com/iris-contrib/middleware/logger"
	"gopkg.in/kataras/iris.v5"

	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
)

func main() {
	// set the favicon
	//iris.Favicon("../frontend/public/images/favicon.ico", "/favicon.ico")

	// set static folder(s)
	//iris.Static("/public", "../frontend/public", 1)

	// set the global middlewares
	iris.Use(logger.New())
	//iris.Use(cors.Default())

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// DB Main
	DbMain()
	// register the routes & the public API
	//registerRoutes()
	registerAPI()

	// start the server
	iris.Listen(":8080")
}

/* func registerRoutes() {
	// register index using a 'Handler'
	iris.Handle("GET", "/", routes.Index())

	// this is other way to declare a route
	// using a 'HandlerFunc'
	iris.Get("/about", routes.About)

	// Dynamic route

	iris.Get("/profile/:username", routes.Profile)("user-profile")
	// user-profile is the custom,optional, route's Name: with this we can use the {{ url "user-profile" $username}} inside userlist.html

	iris.Get("/all", routes.UserList)
} */

func registerAPI() {
	// this is other way to declare routes using the 'API'
	auth := new(api.AuthAPI)
	visitors := new(api.UserAPI)

	// Custom handler
	iris.Handle("GET", "/v1/blog/news", api.CustomAPI{})
	// Function handler
	iris.Post("/v1/auth/login", auth.Login)
	iris.Post("/v1/auth/register", auth.Register)
	iris.Get("/v1/auth/check", auth.Check)
	iris.Get("/v1/auth/session", auth.Session)
	// Api handler
	iris.Get("/v1/visitors", visitors.GetVisitors)
	iris.Get("/v1/visitors/:id", visitors.GetById)
	iris.Put("/v1/visitors/:id", visitors.PutOrPostById)
	iris.Post("/v1/visitors/:id", visitors.Post)
}

func DbMain() {
	Db := db.MgoDb{}
	Db.Init()
	// index keys
	keys := []string{"email", "id", "name"}
	Db.Index("auth", keys)
	Db.Index("people", keys)
}
