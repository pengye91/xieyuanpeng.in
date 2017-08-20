package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/authorization"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/middlewares"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
)

func DbMain() {
	DB := &db.MgoDb{}
	DB.Init()
}

var (
	auth = &api.AuthAPI{}
	pic  = &api.PictureAPI{}
	blog  = &api.BlogAPI{}
)

func main() {
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		fmt.Println(envErr)
	}
	DbMain()
	//gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	go api.InitialUserInRedis()
	go utils.CleanTimeSlice()
	app.Use(middlewares.CORSMiddleware)
	app.Use(middlewares.GlobalStatisticsMiddleware())


	apiV1 := app.Group("/api/v1")
	{
		apiV1.Static("/html", "/home/xyp/go/src/github.com/pengye91/xieyuanpeng.in/static/html")
		apiV1.Static("/md", "/home/xyp/go/src/github.com/pengye91/xieyuanpeng.in/static/md")
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
			p.POST("/:id/comments", middlewares.JWTMiddlewareFactory(authorization.All).MiddlewareFunc(), pic.PostCommentToPic)
			p.PUT("/:id/comments", middlewares.JWTMiddlewareFactory(authorization.All).MiddlewareFunc(), pic.UpdateCommentByPicId)
			p.DELETE("/:id/comments", pic.DeleteCommentByPicId)
		}
		apiV1.POST("/picses", middlewares.JWTMiddlewareFactory(authorization.IsAdmin).MiddlewareFunc(), pic.PostPicsToMain)
		apiV1.POST("/upload-pics", middlewares.JWTMiddlewareFactory(authorization.IsAdmin).MiddlewareFunc(), pic.UploadPicsToStorage)

		b := apiV1.Group("/blogs")
		{
			b.POST("/", blog.PostBlogToMain)
			b.GET("/", blog.GetAllBlogs)
			b.GET("/:id", blog.GetBlogById)
			b.PUT("/:id/like", blog.LikeBlog)
			b.POST("/:id/comments", middlewares.JWTMiddlewareFactory(authorization.All).MiddlewareFunc(), blog.PostCommentToBlog)
			b.PUT("/:id/comments", middlewares.JWTMiddlewareFactory(authorization.All).MiddlewareFunc(), blog.UpdateCommentByBlogId)
			b.DELETE("/:id/comments", blog.DeleteCommentByBlogId)
		}

		u := apiV1.Group("/users")
		{
			u.GET("/auto-search", api.AutoSearch)
		}
	}

	app.Run("0.0.0.0:8000")
}
