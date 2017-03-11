package api

import (
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"

	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"time"
	"strconv"
)

type UserAPI struct {
	*iris.Context
}

// GET /users
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

// GET /users/:param1
func (this UserAPI) GetById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	visitor := models.VisitorBasic{}
	id := ctx.Param("id")

	if err := Db.C("people").Find(bson.M{"id": id}).One(&visitor); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		ctx.JSON(iris.StatusOK, visitor)
	}
	Db.Close()
}

// PUT /users/:param1
func (this UserAPI) PutOrPostById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	visitorInfo := models.VisitorBasic{}
	id := ctx.Param("id")

	if err := ctx.ReadJSON(&visitorInfo); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
		return
	}

	c := Db.C("people")
	colQuerier := bson.M{"id": id}

	// Update
	if count, _ := c.FindId(id).Count(); count != 0 {
		change := bson.M{"$set": bson.M{
			"name":       visitorInfo.Name,
			"email":      visitorInfo.Email,
			"updated_at": time.Now(),
		}}
		err := c.Update(colQuerier, change)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, models.Err("5"))
		} else {
			println(visitorInfo.Name + " has been inserted to database")
			ctx.JSON(iris.StatusOK, iris.Map{"response": true})
		}
	} else {
		// Post
		visitorNumber, _ := Db.C("people").Count()
		println(visitorNumber)
		visitorInfo.Id = strconv.Itoa(visitorNumber + 1)
		visitorInfo.CreatedAt = time.Now()
		visitorInfo.UpdatedAt = time.Now()
		// Insert
		if err := Db.C("people").Insert(&visitorInfo); err != nil {
			ctx.JSON(iris.StatusOK, models.Err("5"))
		} else {
			ctx.JSON(iris.StatusOK, iris.Map{"response": true})
		}
		Db.Close()
	}
}

// POST /users/:param1
func (this UserAPI) Post(ctx *iris.Context) {
	visitorInfo := models.VisitorBasic{}

	if err := ctx.ReadJSON(&visitorInfo); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("4"))
		return
	}
	Db := db.MgoDb{}
	Db.Init()

	visitorNumber, _ := Db.C("people").Count()
	visitorInfo.Id = strconv.Itoa(visitorNumber + 1)
	visitorInfo.CreatedAt = time.Now()
	visitorInfo.UpdatedAt = time.Now()
	// Insert
	if err := Db.C("people").Insert(&visitorInfo); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("5"))
	} else {
		ctx.JSON(iris.StatusOK, &visitorInfo)
	}
	Db.Close()
}

// DELETE /users/:param1
func (this UserAPI) DeleteById(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	id := ctx.Param("id")

	if err := Db.C("people").Remove(bson.M{"id": id}); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("1"))
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true})
	}
}
