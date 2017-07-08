package api

import (
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"
	"time"
	"fmt"
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

func (this PictureAPI) GetPicById(ctx *iris.Context) {
	// TODO: add authentication
	Db := db.MgoDb{}
	Db.Init()
	pic := models.Picture{}

	picId := ctx.Param("id")

	if err := Db.C("picture").FindId(bson.ObjectIdHex(picId)).One(&pic); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, pic)
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
	pic.CreatedAt = time.Now()
	pic.Id = bson.NewObjectId()

	if err := Db.C("picture").Insert(&pic); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	} else {
		ctx.JSON(iris.StatusCreated, pic)
	}
	Db.Close()
}

func (this PictureAPI) PostCommentToPic(ctx *iris.Context) {
	// TODO: a minxin function like login_required()
	Db := db.MgoDb{}
	Db.Init()

	visitorId := ctx.Session().GetString("visitor")

	comment := CommentAlias{}

	picId := ctx.Param("id")

	if err := ctx.ReadJSON(&comment); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("1"))
	} else {
		comment.CommentPreCreateSave(ctx)
		comment.UnderPic = picId

		if commentErr := Db.C("comment").Insert(&comment); commentErr != nil {
			ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
		}

		appendComment := bson.M{
			"$push": bson.M{
				"comments": comment,
			},
		}
		if visitorId != "" {
			if visitorErr := Db.C("people").UpdateId(bson.IsObjectIdHex(visitorId), appendComment); visitorErr != nil {
				ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
			}
		}
		if picErr := Db.C("picture").UpdateId(bson.ObjectIdHex(picId), appendComment); picErr != nil {
			ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
		}
		ctx.JSON(iris.StatusCreated, comment)
	}
	Db.Close()
}

// Delete picture by Id but remain all comments
func (this PictureAPI) DeletePic(ctx *iris.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()
	PicId := ctx.Param("id")

	if err := Db.C("picture").RemoveId(bson.ObjectIdHex(PicId)); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, iris.Map{"details": "deleted"})

	Db.Close()
}


// Delete all pics
func (this PictureAPI) DeletePics(ctx *iris.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()

	if changeInfo, err := Db.C("picture").RemoveAll(nil); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	} else {
		fmt.Printf("%v\n", *changeInfo)
	}
	ctx.JSON(iris.StatusOK, iris.Map{"details": "deleted"})

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

	if err := Db.C("picture").FindId(bson.IsObjectIdHex(picId)).One(&pic); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, pic.Comments)
	Db.Close()
}
