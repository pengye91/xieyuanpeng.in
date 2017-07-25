package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
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

	store, _ := sessions.NewRedisStore(10, "tcp", configs.REDIS_URL, "", []byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   configs.BASE_DOMAIN,
		MaxAge:   86400,
		Secure:   false,
		HttpOnly: false,
	})
	session_middleware := sessions.Sessions("sessionid", store)

	apiV1 := app.Group("/api/v1")
	{
		a := apiV1.Group("/auth", session_middleware)
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
			//p.POST("/:id/comments", pic.PostCommentToPic)
			p.PUT("/:id/comments", middlewares.JWTAuthMiddleware.MiddlewareFunc(), pic.UpdateCommentByPicId)
			p.DELETE("/:id/comments", pic.DeleteCommentByPicId)
		}

		u := apiV1.Group("/users")
		{
			u.GET("/auto-search", api.AutoSearch)
		}

	}

	app.Run("0.0.0.0:8000")
}
