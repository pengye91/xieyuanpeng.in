package api

import (
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/kataras/iris.v5"
	"time"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"strconv"
	"gopkg.in/mgo.v2/bson"
)

type PictureAPI struct {
	*iris.Context
}

func (this PictureAPI) GetAllPics(ctx *iris.Context) {
	// TODO: add authentication
	Db := db.MgoDb{}
	Db.Init()
	pics := []models.Picture{}

	if err := Db.C("picture").Find(nil).All(&pics); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, pics)
	Db.Close()
}

func (this PictureAPI) PostPicToMain(ctx *iris.Context) {
	//TODO: only admin can do this
	Db := db.MgoDb{}
	Db.Init()

	pic := models.Picture{}
	if err := ctx.ReadJSON(&pic); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
	}
	pictureNumber, picCountErr := Db.C("picture").Count()
	if picCountErr != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	pic.Id = strconv.Itoa(pictureNumber + 1)
	pic.CreatedAt = time.Now()

	if err := Db.C("picture").Insert(&pic); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	} else {
		ctx.JSON(iris.StatusCreated, pic)
	}
	Db.Close()
}

func (this PictureAPI) AddCommentToPic(ctx *iris.Context) {
	// TODO: all visitors should can add comment to pic?
	Db := db.MgoDb{}
	Db.Init()

	picId := ctx.Param("id")
	comment := models.Comment{}

	if err := ctx.ReadJSON(&comment); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
	}

	query := bson.M{"id": picId}
	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}
	if err := Db.C("picture").Update(query, update); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, comment)

	Db.Close()
}