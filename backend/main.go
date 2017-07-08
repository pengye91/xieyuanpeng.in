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
)

func main() {
	DbMain()
	app := gin.Default()
	app.Use(cors.Default())

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	session_middleware := sessions.Sessions("sessionid", store)

	a := app.Group("/auth", session_middleware)
	{
		a.POST("/register", auth.Register)
		a.POST("/login", auth.Login)
		a.POST("/check", auth.Check)
	}

	app.Run(":8080")
}
