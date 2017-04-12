package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"gopkg.in/kataras/iris.v5"

	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
)

func main() {
	// set the favicon
	//iris.Favicon("../frontend/public/images/favicon.ico", "/favicon.ico")

	// set static folder(s)
	iris.StaticFS("/static", "../xiyuanpeng_front/public", 1)

	// set the global middlewares
	iris.Use(logger.New())
	iris.Use(cors.Default())

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
	iris.Listen(":8000")
}

func registerAPI() {
	// this is other way to declare routes using the 'API'
	auth := new(api.AuthAPI)
	visitors := new(api.UserAPI)
	comments := new(api.CommentApi)
	pictures := new(api.PictureAPI)

	// Custom handler
	iris.Handle("GET", "/v1/blog/news", api.CustomAPI{})
	// Auth handler
	iris.Post("/v1/auth/login", auth.Login)
	iris.Post("/v1/auth/register", auth.Register)
	iris.Get("/v1/auth/check", auth.Check)
	iris.Get("/v1/auth/session", auth.Session)
	// visitors handler
	iris.Get("/v1/visitors", visitors.GetVisitors)
	iris.Get("/v1/visitors/:id", visitors.GetById)
	iris.Put("/v1/visitors/:id", visitors.PutById)
	iris.Delete("/v1/visitors/:id", visitors.DeleteById)
	// Comment handler
	//iris.Post("/v1/comments", comments.PostCommentToPic)
	iris.Get("/v1/comments", comments.GetAllComments)
	//iris.Put("/v1/comments/:id", comments.PutCommentToPic)
	// Pictures handler
	iris.Post("/v1/pictures", pictures.PostPicToMain)
	iris.Get("/v1/pictures", pictures.GetAllPics)
	iris.Get("/v1/pictures/:id", pictures.GetPicById)
	iris.Delete("/v1/pictures/:id", pictures.DeletePic)
	iris.Post("/v1/pictures/:id/comments", pictures.PostCommentToPic)
	iris.Get("/v1/pictures/:id/comments", pictures.GetPicComments)
}

func DbMain() {
	Db := db.MgoDb{}
	Db.Init()
	// index keys
	keys := []string{"id"}
	Db.Index("auth", keys)
	Db.Index("people", keys)
	Db.Index("picture", keys)
	Db.Index("comment", keys)
}
