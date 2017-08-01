package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"strconv"
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

func (this PictureAPI) PostPicsToMain(ctx *gin.Context) {
	//TODO: only admin can do this
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var insertPics []interface{}
	pics := []models.Picture{}

	if err := ctx.BindJSON(&pics); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	for index, _ := range pics {
		pics[index].CreatedAt = time.Now()
		insertPics = append(insertPics, pics[index])
	}

	// The initialization part should be done in front-end.
	bulk := Db.C("picture").Bulk()
	bulk.Insert(insertPics...)
	if bulkResult, bulkErr := bulk.Run(); bulkErr != nil {
		fmt.Println(bulkErr)
		fmt.Println(bulkResult)
		ctx.JSON(http.StatusInternalServerError, bulkErr)
	} else {
		ctx.JSON(http.StatusCreated, pics)
	}
}

func (this PictureAPI) UploadPicsToStorage(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	files := form.File["pics"]
	contentTypes := form.Value["content-type"]
	sizes := form.Value["size"]
	fmt.Println(form)
	fmt.Println(contentTypes)
	fmt.Println(sizes)

	var intSizes []int64

	for _, v := range sizes {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			intSizes = append(intSizes, int64(tmp))
		}
	}

	//for _, file := range files {
	//	fmt.Println("filename: " + file.Filename)
	//	if err := ctx.SaveUploadedFile(file, filepath.Join(configs.IMAGE_ROOT, file.Filename)); err != nil {
	//		ctx.JSON(http.StatusBadRequest, err)
	//		return
	//	}
	//}

	utils.UploadToS3(files, contentTypes, intSizes)
	ctx.JSON(http.StatusCreated, gin.H{"done": "ok"})
}

func (this PictureAPI) UpdateCommentByPicId(ctx *gin.Context) {
	// TODO: a minxin function like login_required()
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	comment := models.Comment{}

	picId := ctx.Param("id")
	fmt.Printf("%s\n", picId)

	if err := ctx.BindJSON(&comment); err != nil {
		fmt.Println(err)
		return
	} else {
		comment.ModifiedAt = time.Now()
	}

	updateComment := bson.M{
		"$set": bson.M{
			comment.InternalPath + ".word_content": comment.WordContent,
			comment.InternalPath + ".modified_at":  comment.ModifiedAt,
			comment.InternalPath + ".published_at": comment.PublishedAt,
		},
	}

	if picErr := Db.C("picture").UpdateId(bson.ObjectIdHex(picId), updateComment); picErr != nil {
		fmt.Printf("%s\n", picErr)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	ctx.JSON(http.StatusCreated, comment)

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
		ctx.JSON(http.StatusBadRequest, err)
		return
	} else {
		comment.Id = bson.NewObjectId()
		comment.CreatedAt = time.Now()
		comment.ModifiedAt = time.Now()
		comment.PublishedAt = time.Now()
	}

	var appendComment = bson.M{}
	if comment.InternalPath != "" {
		appendComment = bson.M{
			"$push": bson.M{
				comment.InternalPath + ".comments": comment,
			},
		}
	} else {
		appendComment = bson.M{
			"$push": bson.M{
				"comments": comment,
			},
		}
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

func (this PictureAPI) LikePic(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var likedVisitor struct {
		Increase int                  `json:"increase"`
		LikeType string               `json:"likeType"`
		LikedBy  models.VisitorNameId `json:"likedBy" bson:"liked_by"  form:"liked_by"`
	}

	if err := ctx.BindJSON(&likedVisitor); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
		return
	}

	PicId := ctx.Param("id")

	likePic := bson.M{
		"$inc": bson.M{
			"like": likedVisitor.Increase,
		},
		likedVisitor.LikeType: bson.M{
			"liked_by": likedVisitor.LikedBy,
		},
	}

	if picErr := Db.C("picture").UpdateId(bson.ObjectIdHex(PicId), likePic); picErr != nil {
		fmt.Printf("%s\n", picErr)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": "done"})
}
