package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"net/http"
)

func Server(e *echo.Echo) {
	e.Static("/static", "../xiyuanpeng_front/public")
	e.Static("/test", "../xiyuanpeng_front/src/assets")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/test", func(ctx echo.Context) error {
		m := map[string]string{"xixi": "haha"}

		return ctx.JSON(http.StatusOK, m)
	})
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

func main() {
	e := echo.New()
	DbMain()

	Server(e)
	e.Logger.Debug(e.Start(":8000"))
}
