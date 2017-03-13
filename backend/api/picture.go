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

type PictureAlias models.Picture

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
	comment := CommentAlias{}

	if err := ctx.ReadJSON(&comment); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
	}

	comment.CommentPreCreateSave(ctx)
	comment.UnderPic = picId

	picQuery := bson.M{"id": picId}
	update := bson.M{
		"$push": bson.M{
			"comments": &comment,
		},
	}
	if err := Db.C("comment").Insert(&comment); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	// TODO: there should be a visitorQuery to push comments to visitor doc
	if err := Db.C("picture").Update(picQuery, update); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, comment)

	Db.Close()
}

// Delete picture by Id but remain all comments
func (this PictureAPI) DeletePic(ctx *iris.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()
	PicId := ctx.Param("id")

	picture := models.Picture{}
	if err := Db.C("picture").Find(bson.M{"id": PicId}).One(&picture); err != nil {
		ctx.JSON(iris.StatusNotFound, models.Err("5"))
	}
	if err:= Db.C("picture").Remove(bson.M{"id": PicId}); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, picture)

	Db.Close()
}

func (this PictureAPI) UpdatePic(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	Db.Close()
}

func (this PictureAPI) GetPicComments(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()

	picId := ctx.Param("id")

	pic := models.Picture{}

	if err := Db.C("picture").Find(bson.M{"id": picId}).One(&pic); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, pic.Comments)
	Db.Close()
}
