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
)

func main() {
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		fmt.Println(envErr)
	}
	DbMain()
	//gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	go api.InitialSetsFromDB()
	go utils.TestConsumer("nsq_reader", "chan1", utils.OnMessage)
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
			p.POST("/:id/comments", middlewares.JWTMiddlewareFactory(authorization.All).MiddlewareFunc(), pic.PostCommentToPic)
			p.PUT("/:id/comments", middlewares.JWTMiddlewareFactory(authorization.All).MiddlewareFunc(), pic.UpdateCommentByPicId)
			p.DELETE("/:id/comments", pic.DeleteCommentByPicId)
		}
		apiV1.POST("/picses", middlewares.JWTMiddlewareFactory(authorization.IsAdmin).MiddlewareFunc(), pic.PostPicsToMain)
		apiV1.POST("/upload-pics", middlewares.JWTMiddlewareFactory(authorization.IsAdmin).MiddlewareFunc(), pic.UploadPicsToStorage)

		u := apiV1.Group("/users")
		{
			u.GET("/auto-search", api.AutoSearch)
		}

	}

	app.Run("0.0.0.0:8000")
}
