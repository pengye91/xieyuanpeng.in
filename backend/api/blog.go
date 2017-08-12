package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"github.com/pengye91/xieyuanpeng.in/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"path/filepath"
)

type BlogAPI struct {
	*gin.Context
}

type BlogAlias models.Blog

func (this BlogAPI) GetAllPics(ctx *gin.Context) {
	// TODO: add authentication
	Db := db.MgoDb{}
	Db.Init()
	blogs := []models.Blog{}

	if err := Db.C("blog").Find(nil).All(&blogs); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, blogs)
	Db.Close()
}

func (this BlogAPI) GetBlogById(ctx *gin.Context) {
	// TODO: add authentication
	Db := db.MgoDb{}
	Db.Init()
	blog := models.Blog{}

	blogId := ctx.Param("id")

	if err := Db.C("blog").FindId(bson.ObjectIdHex(blogId)).One(&blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, blog)
	Db.Close()
}

func (this BlogAPI) PostBlogToMain(ctx *gin.Context) {
	//TODO: only admin can do this
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	blog := models.Blog{}
	if err := ctx.BindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
	}
	blog.CreatedAt = time.Now()
	blog.Id = bson.NewObjectId()

	if err := Db.C("blog").Insert(&blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	} else {
		ctx.JSON(http.StatusCreated, blog)
	}
}

func (this BlogAPI) PostBlogsToMain(ctx *gin.Context) {
	//TODO: only admin can do this
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var insertBlogs []interface{}
	blogs := []models.Blog{}

	if err := ctx.BindJSON(&blogs); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	for index := range blogs {
		blogs[index].CreatedAt = time.Now()
		insertBlogs = append(insertBlogs, blogs[index])
	}

	// The initialization part should be done in front-end.
	bulk := Db.C("blog").Bulk()
	bulk.Insert(insertBlogs...)
	if bulkResult, bulkErr := bulk.Run(); bulkErr != nil {
		fmt.Println(bulkErr)
		fmt.Println(bulkResult)
		ctx.JSON(http.StatusInternalServerError, bulkErr)
	} else {
		ctx.JSON(http.StatusCreated, blogs)
	}
}

func (this BlogAPI) UploadBlogsToStorage(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	// form.File["blogs"] are passed from front-end.
	// Doing markdown-to-html conversion in js.
	files := form.File["blogs"]

	// should be html
	contentTypes := form.Value["content-type"]
	sizes := form.Value["size"]

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

	if configs.STATIC_S3_STORAGE {
		utils.UploadToS3(files, contentTypes, intSizes)
	} else {
		for _, file := range files {
			fmt.Println("filename: " + file.Filename)
			if err := ctx.SaveUploadedFile(file, filepath.Join(configs.IMAGE_ROOT, file.Filename)); err != nil {
				ctx.JSON(http.StatusBadRequest, err)
				return
			}
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"done": "ok"})
}

func (this BlogAPI) UpdateCommentByBlogId(ctx *gin.Context) {
	// TODO: a minxin function like login_required()
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	comment := models.Comment{}

	blogId := ctx.Param("id")
	fmt.Printf("%s\n", blogId)

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

	if blogErr := Db.C("blog").UpdateId(bson.ObjectIdHex(blogId), updateComment); blogErr != nil {
		fmt.Printf("%s\n", blogErr)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	ctx.JSON(http.StatusCreated, comment)

}
func (this BlogAPI) PostCommentToBlog(ctx *gin.Context) {
	// TODO: a minxin function like login_required()
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	//session := sessions.Default(ctx)

	//visitorId := session.Get("visitor").(string)

	comment := models.Comment{}

	blogId := ctx.Param("id")

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

	if blogErr := Db.C("blog").UpdateId(bson.ObjectIdHex(blogId), appendComment); blogErr != nil {
		fmt.Printf("%s\n", blogErr)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}

// Delete blog by Id but remain all comments
func (this BlogAPI) DeleteBlog(ctx *gin.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()
	BlogId := ctx.Param("id")

	if err := Db.C("blog").RemoveId(bson.ObjectIdHex(BlogId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, gin.H{"details": "deleted"})

	Db.Close()
}

// Delete all Blogs
func (this BlogAPI) DeleteBlogs(ctx *gin.Context) {
	// TODO: add admin authentication
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if changeInfo, err := Db.C("blog").RemoveAll(nil); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	} else {
		fmt.Printf("%v\n", *changeInfo)
	}
	ctx.JSON(http.StatusOK, gin.H{"details": "deleted"})
}

// Blog may need edit, but not pic
func (this BlogAPI) UpdateBlog(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()

	Db.Close()
}

func (this BlogAPI) GetBlogComments(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	blogId := ctx.Param("id")

	blog := models.Blog{}

	if err := Db.C("blog").FindId(bson.IsObjectIdHex(blogId)).One(&blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, blog.Comments)
}

func (this BlogAPI) DeleteCommentByBlogId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	BlogId := ctx.Param("id")
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

	if blogErr := Db.C("blog").UpdateId(bson.ObjectIdHex(BlogId), deleteComment); blogErr != nil {
		fmt.Printf("%s\n", blogErr)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": "done"})
}

func (this BlogAPI) LikeBlog(ctx *gin.Context) {
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

	BlogId := ctx.Param("id")

	likeBlog := bson.M{
		"$inc": bson.M{
			"like": likedVisitor.Increase,
		},
		likedVisitor.LikeType: bson.M{
			"liked_by": likedVisitor.LikedBy,
		},
	}

	if blogErr := Db.C("blog").UpdateId(bson.ObjectIdHex(BlogId), likeBlog); blogErr != nil {
		fmt.Printf("%s\n", blogErr)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": "done"})
}
