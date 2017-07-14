package api

import (
	"fmt"
	"net/http"
	"time"

	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/mgo.v2/bson"
)

type PictureAPI struct {
	*gin.Context
}

type PictureAlias models.Picture

func (this PictureAPI) GetAllPics(ctx *gin.Context) {
	// TODO: add authentication
	Db := db.MgoDb{}
	Db.Init()
	pics := []models.Picture{}

	if err := Db.C("picture").Find(nil).All(&pics); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, pics)
	Db.Close()
}

func (this PictureAPI) GetPicById(ctx *gin.Context) {
	// TODO: add authentication
	Db := db.MgoDb{}
	Db.Init()
	pic := models.Picture{}

	picId := ctx.Param("id")

	if err := Db.C("picture").FindId(bson.ObjectIdHex(picId)).One(&pic); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, pic)
	Db.Close()
}

func (this PictureAPI) PostPicToMain(ctx *gin.Context) {
	//TODO: only admin can do this
	Db := db.MgoDb{}
	Db.Init()

	pic := models.Picture{}
	if err := ctx.BindJSON(&pic); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
	}
	pic.CreatedAt = time.Now()
	pic.Id = bson.NewObjectId()

	if err := Db.C("picture").Insert(&pic); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	} else {
		ctx.JSON(http.StatusCreated, pic)
	}
	Db.Close()
}

func (this PictureAPI) PostCommentToPic(ctx *gin.Context) {
	// TODO: a minxin function like login_required()
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	//session := sessions.Default(ctx)

	//visitorId := session.Get("visitor").(string)

	comment := models.Comment{}

	picId := ctx.Param("id")
	fmt.Printf("%s\n", picId)

	if err := ctx.BindJSON(&comment); err != nil {
		fmt.Println(err)
		return
	} else {
		comment.Id = bson.NewObjectId()
		comment.CreatedAt = time.Now()
		comment.ModifiedAt = time.Now()
		comment.PublishedAt = time.Now()
	}

	appendComment := bson.M{
		"$push": bson.M{
			comment.InternalPath: comment,
		},
	}

	if picErr := Db.C("picture").UpdateId(bson.ObjectIdHex(picId), appendComment); picErr != nil {
		fmt.Printf("%s\n", picErr)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	ctx.JSON(http.StatusCreated, comment)

}

// Delete picture by Id but remain all comments
func (this PictureAPI) DeletePic(ctx *gin.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()
	PicId := ctx.Param("id")

	if err := Db.C("picture").RemoveId(bson.ObjectIdHex(PicId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, gin.H{"details": "deleted"})

	Db.Close()
}

// Delete all pics
func (this PictureAPI) DeletePics(ctx *gin.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()

	if changeInfo, err := Db.C("picture").RemoveAll(nil); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	} else {
		fmt.Printf("%v\n", *changeInfo)
	}
	ctx.JSON(http.StatusOK, gin.H{"details": "deleted"})

	Db.Close()
}

func (this PictureAPI) UpdatePic(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()

	Db.Close()
}

func (this PictureAPI) GetPicComments(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()

	picId := ctx.Param("id")

	pic := models.Picture{}

	if err := Db.C("picture").FindId(bson.IsObjectIdHex(picId)).One(&pic); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, pic.Comments)
	Db.Close()
}

func (this PictureAPI) DeleteCommentByPicId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	PicId := ctx.Param("id")
	internalPath := ctx.Query("internalPath")
	id := ctx.Query("id")

	fmt.Println(id)
	fmt.Println(internalPath)

	deleteComment := bson.M{
		"$pull": bson.M{
			internalPath: bson.M{
				"_id": bson.ObjectIdHex(id),
			},
		},
	}

	if picErr := Db.C("picture").UpdateId(bson.ObjectIdHex(PicId), deleteComment); picErr != nil {
		fmt.Printf("%s\n", picErr)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": "done"})

}
