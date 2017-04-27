package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	// "github.com/iris-contrib/middleware/jwt"
	"gopkg.in/kataras/iris.v5"

	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
)

var (
	MyCorsConfig *cors.Options
	app          *iris.Framework
	auth         *api.AuthAPI
	visitors     *api.UserAPI
	comments     *api.CommentApi
	pictures     *api.PictureAPI
)

func init() {
	MyCorsConfig = &cors.Options{
		AllowedMethods: []string{"GET", "POST", "HEAD", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		Debug:          true,
		OptionsPassthrough: false,
	}
	auth = new(api.AuthAPI)
	visitors = new(api.UserAPI)
	comments = new(api.CommentApi)
	pictures = new(api.PictureAPI)
}

func main() {
	// set the favicon
	//iris.Favicon("../frontend/public/images/favicon.ico", "/favicon.ico")

	// set static folder(s)

	// set the global middlewares
	app = iris.New()
	app.Use(logger.New())
	//app.Use(cors.New(*MyCorsConfig))

	// set the custom errors
	//app.OnError(app.StatusNotFound, func(ctx *app.Context) {
	//	ctx.Render("errors/404.html", app.Map{"Title": app.StatusText(app.StatusNotFound)})
	//})

	//app.OnError(app.StatusInternalServerError, func(ctx *app.Context) {
	//	ctx.Render("errors/500.html", nil, app.RenderOptions{"layout": app.NoLayout})
	//})
	// DB Main
	DbMain()

	a := app.Party("/v1/auth", cors.New(*MyCorsConfig).Serve)
	{
		a.Post("/login", auth.Login)
		a.Post("/register", auth.Register)
		a.Get("/check", auth.Check)
		a.Get("/session", auth.Session)
	}
	v := app.Party("/v1/visitors")
	{
		v.Get("/", visitors.GetVisitors)
		v.Get("/:id", visitors.GetById)
		v.Put("/:id", visitors.PutById)
		v.Delete("/:id", visitors.DeleteById)
	}
	c := app.Party("/v1/comments")
	{
		c.Get("/", comments.GetAllComments)
	}
	p := app.Party("/v1/pictures")
	{
		p.Post("/", pictures.PostPicToMain)
		p.Get("/", pictures.GetAllPics)
		p.Get("/:id", pictures.GetPicById)
		p.Delete("/:id", pictures.DeletePic)
		p.Delete("/", pictures.DeletePics)
		p.Post("/:id/comments", pictures.PostCommentToPic)
		p.Get("/:id/comments", pictures.GetPicComments)
	}

	app.Listen(":8000")
	app.StaticServe("../xiyuanpeng_front/public", "/static")
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
