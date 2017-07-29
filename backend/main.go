package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/middlewares"
)

func DbMain() {
	DB := &db.MgoDb{}
	DB.Init()
}

var (
	auth = &api.AuthAPI{}
	pic  = &api.PictureAPI{}
)

func main() {
	DbMain()
	//gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	go api.InitialSetsFromDB()
	app.Use(middlewares.CORSMiddleware)

	apiV1 := app.Group("/api/v1")
	{
		a := apiV1.Group("/auth", middlewares.Session_middleware)
		{
			a.POST("/register", auth.Register)
			a.POST("/login", middlewares.JWTAuthMiddleware.LoginHandler)
			a.GET("/logout", auth.LogOut)
			a.GET("/refresh_token", middlewares.JWTAuthMiddleware.RefreshHandler)
			a.POST("/check", auth.Check)
			a.GET("/allusers", auth.GetAllUsers)
		}

		p := apiV1.Group("/pics")
		{
			p.POST("/", pic.PostPicToMain)
			p.GET("/", pic.GetAllPics)
			p.GET("/:id", pic.GetPicById)
			p.PUT("/:id/like", pic.LikePic)
			p.POST("/:id/comments", middlewares.JWTAuthMiddleware.MiddlewareFunc(), pic.PostCommentToPic)
			p.PUT("/:id/comments", middlewares.JWTAuthMiddleware.MiddlewareFunc(), pic.UpdateCommentByPicId)
			p.DELETE("/:id/comments", pic.DeleteCommentByPicId)
		}
		apiV1.POST("/picses", pic.PostPicsToMain)

		u := apiV1.Group("/users")
		{
			u.GET("/auto-search", api.AutoSearch)
		}

	}

	app.Run("0.0.0.0:8000")
}
