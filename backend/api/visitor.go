package api

import (
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"time"
)

type UserAPI struct {
	*iris.Context
}

func (this UserAPI) GetVisitors(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	visitors := []models.Visitor{}
	if err := Db.C("people").Find(nil).All(&visitors); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		ctx.JSON(iris.StatusOK, &visitors)
	}

	Db.Close()
}

func (this UserAPI) GetById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	visitor := models.Visitor{}
	id := ctx.Param("id")

	if err := Db.C("people").Find(bson.M{"id": id}).One(&visitor); err != nil {
		ctx.JSON(iris.StatusNotFound, models.Err("1"))
		return
	} else {
		ctx.JSON(iris.StatusOK, visitor)
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
	colQuerier := bson.M{"id": id}

	// Update
	if count, countErr := c.Find(bson.M{"id": id}).Count(); count != 0 {
		change := bson.M{"$set": bson.M{
				"basic.name":       visitorInfo.Name,
				"basic.email":      visitorInfo.Email,
				"basic.updated_at": time.Now(),
			}}
		err := c.Update(colQuerier, change)
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

	if err := Db.C("people").Remove(bson.M{"id": id}); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("1"))
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": "successfully delete"})
	}
}
