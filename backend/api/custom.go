package api

import (
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/kataras/iris.v5"
)

type CustomAPI struct {
	*iris.Context
}

func (this CustomAPI) Serve(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	results := []models.Visitor{}

	if err := Db.C("auth").Find(nil).All(&results); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("5"))
		return
	}

	ctx.JSON(iris.StatusOK, &results)
	Db.Close()
}
