package api

import (
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
	"time"
)

type UserAPI struct {
	*iris.Context
}

var (
	logger, _    = utils.MyDevLogger()
	errLogger, _ = utils.MyErrLogger()
	sugar        = logger.Sugar()
	errSugar = errLogger.Sugar()
)

func (this UserAPI) GetVisitors(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	visitors := []models.Visitor{}
	if err := Db.C("people").Find(nil).All(&visitors); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		ctx.JSON(iris.StatusOK, &visitors)
		sugar.Infow(
			"200 OK",
			"url", ctx.PathString(),
		)
	}

	Db.Close()
}

func (this UserAPI) GetById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	visitor := models.Visitor{}
	id := ctx.Param("id")
	if !bson.IsObjectIdHex(ctx.Param("id")) {
		//errSugar.Errorw(
		//	"500 Internal Server Error ",
		//	"url", ctx.URI(),
		//)
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
	} else {
		if err := Db.C("people").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&visitor); err != nil {
			sugar.Warnw(
				"404 Not Found",
				"url", ctx.URI(),
			)
			ctx.JSON(iris.StatusNotFound, models.Err("1"))
			return
		} else {
			ctx.JSON(iris.StatusOK, visitor)
		}
	}
	Db.Close()
}

func (this UserAPI) PutById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	visitorInfo := models.VisitorBasic{}
	id := ctx.Param("id")

	if err := ctx.ReadJSON(&visitorInfo); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
		return
	} else {
		fmt.Printf("%v\n", visitorInfo)
	}

	c := Db.C("people")

	// Update
	if count, countErr := c.FindId(bson.IsObjectIdHex(id)).Count(); count != 0 {
		change := bson.M{"$set": bson.M{
			"basic.name":       visitorInfo.Name,
			"basic.email":      visitorInfo.Email,
			"basic.updated_at": time.Now(),
		}}
		err := c.UpdateId(bson.IsObjectIdHex(id), change)
		if err != nil {
			//ctx.JSON(iris.StatusBadRequest, models.Err("5"))
			panic(err)
		} else {
			println(visitorInfo.Name + " has been inserted to database")
			ctx.JSON(iris.StatusOK, visitorInfo)
		}
	} else {
		//ctx.JSON(iris.StatusBadRequest, models.Err("5"))
		panic(countErr)
	}
}

// DELETE /users/:param1
func (this UserAPI) DeleteById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	id := ctx.Param("id")

	if err := Db.C("people").RemoveId(bson.IsObjectIdHex(id)); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("1"))
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": "successfully delete"})
	}
}
