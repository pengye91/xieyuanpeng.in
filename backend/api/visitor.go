package api
<<<<<<< HEAD
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
=======

//import (
//	"gopkg.in/kataras/iris.v5"
//	"gopkg.in/mgo.v2/bson"
//
//	"fmt"
//	"github.com/pengye91/xieyuanpeng.in/backend/db"
//	"github.com/pengye91/xieyuanpeng.in/backend/models"
//	"github.com/pengye91/xieyuanpeng.in/backend/utils"
//	"time"
//)
//
//type UserAPI struct {
//	*iris.Context
>>>>>>> dev
//}
//
//var (
//	logger, _    = utils.MyDevLogger()
//	errLogger, _ = utils.MyErrLogger()
//	sugar        = logger.Sugar()
//	errSugar = errLogger.Sugar()
//)
//
<<<<<<< HEAD
//func (this UserAPI) GetVisitors(ctx *gin.Context) {
=======
//func (this UserAPI) GetVisitors(ctx *iris.Context) {
>>>>>>> dev
//	Db := db.MgoDb{}
//	Db.Init()
//
//	visitors := []models.Visitor{}
//	if err := Db.C("people").Find(nil).All(&visitors); err != nil {
<<<<<<< HEAD
//		ctx.JSON(http.StatusOK, models.Err("1"))
//		return
//	} else {
//		ctx.JSON(http.StatusOK, &visitors)
//		sugar.Infow(
//			"200 OK",
//			"url", ctx.Request.URL.Path,
=======
//		ctx.JSON(iris.StatusOK, models.Err("1"))
//		return
//	} else {
//		ctx.JSON(iris.StatusOK, &visitors)
//		sugar.Infow(
//			"200 OK",
//			"url", ctx.PathString(),
>>>>>>> dev
//		)
//	}
//
//	Db.Close()
//}
//
<<<<<<< HEAD
//func (this UserAPI) GetById(ctx *gin.Context) {
=======
//func (this UserAPI) GetById(ctx *iris.Context) {
>>>>>>> dev
//	Db := db.MgoDb{}
//	Db.Init()
//	visitor := models.Visitor{}
//	id := ctx.Param("id")
//	if !bson.IsObjectIdHex(ctx.Param("id")) {
//		//errSugar.Errorw(
//		//	"500 Internal Server Error ",
//		//	"url", ctx.URI(),
//		//)
<<<<<<< HEAD
//		ctx.JSON(http.StatusBadRequest, models.Err("5"))
=======
//		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
>>>>>>> dev
//	} else {
//		if err := Db.C("people").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&visitor); err != nil {
//			sugar.Warnw(
//				"404 Not Found",
<<<<<<< HEAD
//				"url", ctx.Request.URL.Path,
//			)
//			ctx.JSON(http.StatusNotFound, models.Err("1"))
//			return
//		} else {
//			ctx.JSON(http.StatusOK, visitor)
=======
//				"url", ctx.URI(),
//			)
//			ctx.JSON(iris.StatusNotFound, models.Err("1"))
//			return
//		} else {
//			ctx.JSON(iris.StatusOK, visitor)
>>>>>>> dev
//		}
//	}
//	Db.Close()
//}
//
<<<<<<< HEAD
//func (this UserAPI) PutById(ctx *gin.Context) {
=======
//func (this UserAPI) PutById(ctx *iris.Context) {
>>>>>>> dev
//	Db := db.MgoDb{}
//	Db.Init()
//
//	visitorInfo := models.VisitorBasic{}
//	id := ctx.Param("id")
//
<<<<<<< HEAD
//	if err := ctx.BindJSON(&visitorInfo); err != nil {
//		ctx.JSON(http.StatusBadRequest, models.Err("5"))
=======
//	if err := ctx.ReadJSON(&visitorInfo); err != nil {
//		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
>>>>>>> dev
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
<<<<<<< HEAD
//			//ctx.JSON(http.StatusBadRequest, models.Err("5"))
//			panic(err)
//		} else {
//			println(visitorInfo.Name + " has been inserted to database")
//			ctx.JSON(http.StatusOK, visitorInfo)
//		}
//	} else {
//		//ctx.JSON(http.StatusBadRequest, models.Err("5"))
=======
//			//ctx.JSON(iris.StatusBadRequest, models.Err("5"))
//			panic(err)
//		} else {
//			println(visitorInfo.Name + " has been inserted to database")
//			ctx.JSON(iris.StatusOK, visitorInfo)
//		}
//	} else {
//		//ctx.JSON(iris.StatusBadRequest, models.Err("5"))
>>>>>>> dev
//		panic(countErr)
//	}
//}
//
//// DELETE /users/:param1
<<<<<<< HEAD
//func (this UserAPI) DeleteById(ctx *gin.Context) {
=======
//func (this UserAPI) DeleteById(ctx *iris.Context) {
>>>>>>> dev
//	Db := db.MgoDb{}
//	Db.Init()
//
//	id := ctx.Param("id")
//
//	if err := Db.C("people").RemoveId(bson.IsObjectIdHex(id)); err != nil {
<<<<<<< HEAD
//		ctx.JSON(http.StatusBadRequest, models.Err("1"))
//	} else {
//		ctx.JSON(http.StatusOK, gin.H{"response": "successfully delete"})
=======
//		ctx.JSON(iris.StatusBadRequest, models.Err("1"))
//	} else {
//		ctx.JSON(iris.StatusOK, iris.Map{"response": "successfully delete"})
>>>>>>> dev
//	}
//}
