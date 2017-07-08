package api

import (
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type CommentApi struct {
	*iris.Context
}

type CommentAlias models.Comment

func (this CommentApi) GetAllComments(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	comments := []CommentAlias{}
	if err := Db.C("comment").Find(nil).All(&comments); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) GetAllCommentsByPicId(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	picId := ctx.Param("id")
	comments := []CommentAlias{}

	err := Db.C("comment").FindId(bson.M{"UnderPic": picId}).All(&comments)
	if err != nil {
		ctx.JSON(iris.StatusNotFound, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) GetAllCommentsByVisitorId(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	visitorId := ctx.Param("id")
	comments := []CommentAlias{}

	err := Db.C("comment").Find(bson.M{"ById": visitorId}).All(&comments)
	if err != nil {
		ctx.JSON(iris.StatusNotFound, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) PutCommentToPic(ctx *iris.Context) {
	// TODO: authentication
	id := ctx.Param("id")
	comment := models.Comment{}
	if err := ctx.ReadJSON(&comment); err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
	}
	Db := db.MgoDb{}
	Db.Init()
	visitor := models.Visitor{}

	query := bson.M{"comments.id": id}
	update := bson.M{
		"$set": bson.M{
			"comments.$.word_content":     comment.WordContent,
			"comments.$.contain_pic_path": comment.ContainPicPath,
			"comments.$.published_at":     time.Now(),
			"comments.$.modified_at":      time.Now(),
		},
	}

	if err := Db.C("people").Update(query, update); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	} else {
		Db.C("people").Find(bson.M{"comments.id": id}).One(&visitor)
		ctx.JSON(iris.StatusOK, visitor.Comments)
	}
	Db.Close()
}

func (this CommentApi) DeleteComment(ctx *iris.Context) {

}

func (comment *CommentAlias) CommentPreCreateSave(ctx *iris.Context) {
	// TODO: add a global chan to count comment number and other number
	Db := db.MgoDb{}
	Db.Init()

	visitorId := ctx.Session().GetString("visitor")

	comment.Id = bson.NewObjectId()
	if visitorId != "" {
		comment.ById = visitorId
	}
	comment.CreatedAt = time.Now()
	comment.PublishedAt = time.Now()
	comment.ModifiedAt = time.Now()

	Db.Close()
}
