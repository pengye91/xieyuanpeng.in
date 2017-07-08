package api

//import (
//	"net/http"
//	"time"
//
//	"github.com/gin-contrib/sessions"
//	"github.com/gin-gonic/gin"
//	"github.com/pengye91/xieyuanpeng.in/backend/db"
//	"github.com/pengye91/xieyuanpeng.in/backend/models"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type CommentApi struct {
//	*gin.Context
//}
//
//type CommentAlias models.Comment
//
//func (this CommentApi) GetAllComments(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//	comments := []CommentAlias{}
//	if err := Db.C("comment").Find(nil).All(&comments); err != nil {
//		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
//	}
//	ctx.JSON(http.StatusOK, comments)
//	Db.Close()
//}
//
//func (this CommentApi) GetAllCommentsByPicId(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//	picId := ctx.Param("id")
//	comments := []CommentAlias{}
//
//	err := Db.C("comment").FindId(bson.M{"UnderPic": picId}).All(&comments)
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, models.Err("5"))
//	}
//	ctx.JSON(http.StatusOK, comments)
//	Db.Close()
//}
//
//func (this CommentApi) GetAllCommentsByVisitorId(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//	visitorId := ctx.Param("id")
//	comments := []CommentAlias{}
//
//	err := Db.C("comment").Find(bson.M{"ById": visitorId}).All(&comments)
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, models.Err("5"))
//	}
//	ctx.JSON(http.StatusOK, comments)
//	Db.Close()
//}
//
//func (this CommentApi) PutCommentToPic(ctx *gin.Context) {
//	// TODO: authentication
//	id := ctx.Param("id")
//	comment := models.Comment{}
//	if err := ctx.BindJSON(&comment); err != nil {
//		ctx.JSON(http.StatusBadRequest, models.Err("5"))
//	}
//	Db := db.MgoDb{}
//	Db.Init()
//	visitor := models.Visitor{}
//
//	query := bson.M{"comments.id": id}
//	update := bson.M{
//		"$set": bson.M{
//			"comments.$.word_content":     comment.WordContent,
//			"comments.$.contain_pic_path": comment.ContainPicPath,
//			"comments.$.published_at":     time.Now(),
//			"comments.$.modified_at":      time.Now(),
//		},
//	}
//
//	if err := Db.C("people").Update(query, update); err != nil {
//		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
//	} else {
//		Db.C("people").Find(bson.M{"comments.id": id}).One(&visitor)
//		ctx.JSON(http.StatusOK, visitor.Comments)
//	}
//	Db.Close()
//}
//
//func (this CommentApi) DeleteComment(ctx *gin.Context) {
//
//}
//
//func (comment *CommentAlias) CommentPreCreateSave(ctx *gin.Context) {
//	// TODO: add a global chan to count comment number and other number
//	Db := db.MgoDb{}
//	Db.Init()
//
//	session := sessions.Default(ctx)
//
//	visitorId := session.Get("visitorId").(string)
//
//	comment.Id = bson.NewObjectId()
//	if visitorId != "" {
//		comment.ById = visitorId
//	}
//	comment.CreatedAt = time.Now()
//	comment.PublishedAt = time.Now()
//	comment.ModifiedAt = time.Now()
//
//	Db.Close()
//}
