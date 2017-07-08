package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.Logger())

	ee := e.Group("/v1", middleware.CORS())
	{
		ee.GET("/auth/get/*", func(c echo.Context) error {
			if c.QueryParam("name") != "xixi" {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			return c.String(http.StatusOK, "xixi")
		})
		ee.GET("/auth/post", func(c echo.Context) error {
			return c.String(http.StatusOK, "haha")
		})
	}
	e.Logger.Info(e.Start(":1213"))
}
