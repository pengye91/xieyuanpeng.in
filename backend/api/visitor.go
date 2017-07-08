package api
//
//import (
//	"fmt"
//	"net/http"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	"github.com/pengye91/xieyuanpeng.in/backend/db"
//	"github.com/pengye91/xieyuanpeng.in/backend/models"
//	"github.com/pengye91/xieyuanpeng.in/backend/utils"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type UserAPI struct {
//	*gin.Context
//}
//
//var (
//	logger, _    = utils.MyDevLogger()
//	errLogger, _ = utils.MyErrLogger()
//	sugar        = logger.Sugar()
//	errSugar = errLogger.Sugar()
//)
//
//func (this UserAPI) GetVisitors(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//
//	visitors := []models.Visitor{}
//	if err := Db.C("people").Find(nil).All(&visitors); err != nil {
//		ctx.JSON(http.StatusOK, models.Err("1"))
//		return
//	} else {
//		ctx.JSON(http.StatusOK, &visitors)
//		sugar.Infow(
//			"200 OK",
//			"url", ctx.Request.URL.Path,
//		)
//	}
//
//	Db.Close()
//}
//
//func (this UserAPI) GetById(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//	visitor := models.Visitor{}
//	id := ctx.Param("id")
//	if !bson.IsObjectIdHex(ctx.Param("id")) {
//		//errSugar.Errorw(
//		//	"500 Internal Server Error ",
//		//	"url", ctx.URI(),
//		//)
//		ctx.JSON(http.StatusBadRequest, models.Err("5"))
//	} else {
//		if err := Db.C("people").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&visitor); err != nil {
//			sugar.Warnw(
//				"404 Not Found",
//				"url", ctx.Request.URL.Path,
//			)
//			ctx.JSON(http.StatusNotFound, models.Err("1"))
//			return
//		} else {
//			ctx.JSON(http.StatusOK, visitor)
//		}
//	}
//	Db.Close()
//}
//
//func (this UserAPI) PutById(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//
//	visitorInfo := models.VisitorBasic{}
//	id := ctx.Param("id")
//
//	if err := ctx.BindJSON(&visitorInfo); err != nil {
//		ctx.JSON(http.StatusBadRequest, models.Err("5"))
//		return
//	} else {
//		fmt.Printf("%v\n", visitorInfo)
//	}
//
//	c := Db.C("people")
//
//	// Update
//	if count, countErr := c.FindId(bson.IsObjectIdHex(id)).Count(); count != 0 {
//		change := bson.M{"$set": bson.M{
//			"basic.name":       visitorInfo.Name,
//			"basic.email":      visitorInfo.Email,
//			"basic.updated_at": time.Now(),
//		}}
//		err := c.UpdateId(bson.IsObjectIdHex(id), change)
//		if err != nil {
//			//ctx.JSON(http.StatusBadRequest, models.Err("5"))
//			panic(err)
//		} else {
//			println(visitorInfo.Name + " has been inserted to database")
//			ctx.JSON(http.StatusOK, visitorInfo)
//		}
//	} else {
//		//ctx.JSON(http.StatusBadRequest, models.Err("5"))
//		panic(countErr)
//	}
//}
//
//// DELETE /users/:param1
//func (this UserAPI) DeleteById(ctx *gin.Context) {
//	Db := db.MgoDb{}
//	Db.Init()
//
//	id := ctx.Param("id")
//
//	if err := Db.C("people").RemoveId(bson.IsObjectIdHex(id)); err != nil {
//		ctx.JSON(http.StatusBadRequest, models.Err("1"))
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{"response": "successfully delete"})
//	}
//}
