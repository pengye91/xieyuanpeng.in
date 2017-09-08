package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pengye91/xieyuanpeng.in/backend/api"
	"github.com/pengye91/xieyuanpeng.in/backend/authorization"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/middlewares"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/background"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
)

var (
	auth = &api.AuthAPI{}
	pic  = &api.PictureAPI{}
	blog = &api.BlogAPI{}
	menu = &api.MenuApi{}
)

func DbMain() {
	DB := &db.MgoDb{}
	DB.Init()
}

func init() {
	Db := db.MgoDb{}
	Db.Init()
	Db.Index(
		"auth",
		[]string{"name", "email"},
	)
	Db.Index(
		"picture",
		[]string{"title"},
	)
	Db.Index(
		"blog",
		[]string{"title"},
	)

	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.LoggerSugar.Errorw("godotenv.Load file Error",
			"module", "godotenv",
			"error", envErr,
		)
	}
}

func main() {
	DbMain()
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middlewares.GlobalLoggingMiddleware())

	go api.InitialUserInRedis()
	go background.CleanTimeSlice()

	// this two function only need to run one time.
	//go utils.ImportCitiesToRedis("/home/xyp/go/src/github.com/pengye91/xieyuanpeng.in/backend/utils/ip_scripts/GeoLite2-City-CSV_20170801/GeoLite2-City-Locations-zh-CN.csv")
	//go utils.ImportIPToRedis("/home/xyp/go/src/github.com/pengye91/xieyuanpeng.in/backend/utils/ip_scripts/GeoLite2-City-CSV_20170801/GeoLite2-City-Blocks-IPv4.csv")
	app.Use(middlewares.CORSMiddleware)
	app.Use(middlewares.GlobalStatisticsMiddleware())

	//if cityInfo, err := cache.FindCityByIP("110.185.16.73"); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(cityInfo)
	//}

	//for i := 0; i < 20; i++ {
	//	if id, err := sync.AcquireFairSemaphore("testSema", 10, 30*time.Second); err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(id)
	//	}
	//}
	//utils.ReleaseSemaphoreBasedOnTime("testSema", "")

	apiV1 := app.Group("/api/v1")
	{
		m := apiV1.Group("/menu")
		{
			m.GET("/", menu.GetMenu)
			// post hostname/api/v1/menu/ allMenuItems
			m.POST("/", menu.PostMenu)
			// PUT hostname/api/v1/menu/ just one menuItem
			m.PUT("/", menu.PutMenuItem)
			m.GET("/side-menu", menu.GetSideMenu)
			// PUT hostname/api/v1/menu/side-menu/ all sideMenuItems
			m.PUT("/side-menu", menu.PutSideMenuItem)
			m.GET("/admin-side-menu", menu.GetAdminSideMenu)
			// PUT hostname/api/v1/menu/admin-side-menu/ all adminSideMenuItems
			m.PUT("/admin-side-menu", menu.PutAdminSideMenuItem)
		}

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
			b.DELETE("/", blog.DeleteBlogs)
		}
		apiV1.POST("/blogses", middlewares.JWTMiddlewareFactory(authorization.IsAdmin).MiddlewareFunc(), blog.PostBlogsToMain)
		apiV1.POST("/upload-blogs", middlewares.JWTMiddlewareFactory(authorization.IsAdmin).MiddlewareFunc(), blog.UploadBlogsToStorage)

		u := apiV1.Group("/users")
		{
			u.GET("/auto-search", api.AutoSearch)
		}

	}

	app.Run("0.0.0.0:8000")
}
