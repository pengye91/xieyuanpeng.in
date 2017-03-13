package api

import (
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type CommentApi struct {
	*iris.Context
}

type CommentAlias models.Comment

func (this CommentApi) GetAllComments(ctx *iris.Context) {
	Db := db.MgoDb{}
	Db.Init()
	comments := []models.Comment{}
	if err := Db.C("comment").Find(nil).All(&comments); err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(iris.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) PostComment(ctx *iris.Context) {
	// TODO: a minxin function like login_required()

	Db := db.MgoDb{}
	Db.Init()
	if ctx.Session().GetString("login") == "true" {
		comment := models.Comment{}
		if err := ctx.ReadJSON(&comment); err != nil {
			ctx.JSON(iris.StatusBadRequest, models.Err("1"))
		} else {
			visistorId := ctx.Session().GetString("visitor")
			visitor := models.Visitor{}
			Db.C("people").Find(bson.M{"id": visistorId}).One(&visitor)
			commentId := strconv.Itoa(len(visitor.Comments) + 1)
			comment.Id = commentId
			comment.ById = visistorId
			comment.CreatedAt = time.Now()
			comment.PublishedAt = time.Now()
			comment.ModifiedAt = time.Now()

			query := bson.M{"id": visistorId}
			appendComment := bson.M{
				"$push": bson.M{
					"comments": comment,
				},
			}
			err := Db.C("people").Update(query, appendComment)
			if err != nil {
				ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
			}
			ctx.JSON(iris.StatusCreated, comment)
		}
	} else {
		ctx.JSON(iris.StatusForbidden, iris.Map{"detail": "you should login to post a comment."})
	}
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
			"comments.$.word_content": comment.WordContent,
			"comments.$.contain_pic_path": comment.ContainPicPath,
			"comments.$.published_at": time.Now(),
			"comments.$.modified_at": time.Now(),
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

func (comment *CommentAlias) PreCreateSave(ctx *iris.Context) {
	// TODO: add a global chan to count comment number and other number
	Db := db.MgoDb{}
	Db.Init()

	commentNumber, err := Db.C("comment").Count()
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, models.Err("5"))
	}
	comment.Id = strconv.Itoa(commentNumber + 1)
	comment.CreatedAt = time.Now()
	comment.PublishedAt = time.Now()
	comment.ModifiedAt = time.Now()

	Db.Close()
}
