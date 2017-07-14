package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
)

func DbMain() {
	DB := &db.MgoDb{}
	DB.Init()
}

var (
	auth = &api.AuthAPI{}
	pic = &api.PictureAPI{}
	com = &api.CommentApi{}

)

func main() {
	DbMain()
	app := gin.Default()
	xypConfig := cors.DefaultConfig()
	xypConfig.AllowMethods = append(xypConfig.AllowMethods, "DELETE")
	//xypConfig.AllowAllOrigins = true
	xypConfig.AllowOrigins = []string{"http://localhost:8080"}
	xypConfig.AllowCredentials = true
	//xypConfig.AllowHeaders = []string{"*"}
	corsMiddleware := cors.New(xypConfig)
	app.Use(corsMiddleware)

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{
		Path: "/",
		Domain: "localhost",
	})
	session_middleware := sessions.Sessions("sessionid", store)

	a := app.Group("/auth", session_middleware)
	{
		a.POST("/register", auth.Register)
		a.POST("/login", auth.Login)
		a.POST("/check", auth.Check)
		a.GET("/allusers", auth.GetAllUsers)
	}

	p := app.Group("/pics", session_middleware)
	{
		p.POST("/", pic.PostPicToMain)
		p.GET("/", pic.GetAllPics)
		p.GET("/:id", pic.GetPicById)
		p.POST("/:id/comments", pic.PostCommentToPic)
		p.POST("/:id/responses", com.PostCommentToCommentByPicId)
		p.DELETE("/:id/comments", pic.DeleteCommentByPicId)
	}

	c := app.Group("/coms", session_middleware)
	{
		c.GET("/:id/responses", com.GetAllResponsesByCommentId)
		c.POST("/:id/responses", com.PostResponsesByCommentId)
	}

	r := app.Group("/resps", session_middleware)
	{
		r.POST("/:id/responses", com.PostResponsesToResponseByResponseId)
	}
	app.Run(":8000")
}
