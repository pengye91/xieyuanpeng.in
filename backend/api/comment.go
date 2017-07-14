package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/mgo.v2/bson"
)

// only records two levels, from 3rd and on just replace all nested responses.
type CommentApi struct {
	*gin.Context
}

type CommentAlias models.Comment

func (this CommentApi) GetAllComments(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	comments := []CommentAlias{}
	if err := Db.C("comment").Find(nil).All(&comments); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) GetAllCommentsByPicId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	picId := ctx.Param("id")
	comments := []CommentAlias{}

	err := Db.C("comment").FindId(bson.M{"UnderPic": picId}).All(&comments)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) GetAllCommentsByVisitorId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	visitorId := ctx.Param("id")
	comments := []CommentAlias{}

	err := Db.C("comment").Find(bson.M{"ById": visitorId}).All(&comments)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Err("5"))
	}
	ctx.JSON(http.StatusOK, comments)
	Db.Close()
}

func (this CommentApi) GetAllResponsesByCommentId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	CommentId := ctx.Param("id")
	pipe := Db.C("picture").Pipe(
		// very tricky
		[]bson.M{
			{
				// find corresponding picture
				"$match": bson.M{
					"comments._id": bson.ObjectIdHex(CommentId),
				},
			},
			{
				"$unwind": "$comments",
			},
			{
				"$project": bson.M{
					"commentId":        "$comments._id",
					"commentResponses": "$comments.responses",
					"wordContent":      "$comments.word_content",
				},
			},
			{
				// find the corresponding comments
				"$match": bson.M{
					"commentId": bson.ObjectIdHex(CommentId),
				},
			},

		},
	)

	result := bson.M{}
	if err := pipe.One(&result); err != nil {
		fmt.Printf("%s", err)
		return
	}

	ctx.JSON(http.StatusOK, result)
	Db.Close()
}

func (this CommentApi) PostResponsesByCommentId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	CommentId := ctx.Param("id")
	response := models.Comment{}

	if bindJsonErr := ctx.BindJSON(&response); bindJsonErr != nil {
		fmt.Printf("%s\n", bindJsonErr)
	} else {
		response.Id = bson.NewObjectId()
		response.CreatedAt = time.Now()
		response.PublishedAt = time.Now()
		response.ModifiedAt = time.Now()
	}

	query := bson.M{
		"comments._id": bson.ObjectIdHex(CommentId),
	}
	appendResponse := bson.M{
		"$push": bson.M{
			"comments.$.responses": response,
		},
	}

	if picErr := Db.C("picture").Update(query, appendResponse); picErr != nil {
		fmt.Printf("%s\n", picErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	Db.Close()
}

func (this CommentApi) DeleteCommentByPicId(ctx *gin.Context) {
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

func (this CommentApi) PostCommentToCommentByPicId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	PicId := ctx.Param("id")
	response := models.Comment{}

	if bindJsonErr := ctx.BindJSON(&response); bindJsonErr != nil {
		fmt.Printf("%s\n", bindJsonErr)
	} else {
		response.Id = bson.NewObjectId()
		response.CreatedAt = time.Now()
		response.PublishedAt = time.Now()
		response.ModifiedAt = time.Now()
	}

	appendResponse := bson.M{
		"$push": bson.M{
			response.InternalPath: response,
		},
	}

	if picErr := Db.C("picture").UpdateId(bson.ObjectIdHex(PicId), appendResponse); picErr != nil {
		fmt.Printf("%s\n", picErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Since Mongodb cannot use $push to update nested arrays in arbitrary depth, just replace all
// deeply nested arrays using pre-copied comments, which should be down in Gin instead of Front-end
// because of the data size.
func (this CommentApi) PostResponsesToResponseByResponseId(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	ResponseId := ctx.Param("id")
	response := models.Comment{}
	//type firstResponse struct {
	//	Id                 bson.ObjectId `json:"id" bson:"_id"  form:"id"`
	//	FirstResponseId    bson.ObjectId `json:"firstResponseId" bson:"firstResponseId"  form:"first_response_id"`
	//	CommentId          bson.ObjectId `json:"commentId" bson:"commentId"  form:"commentId"`
	//	WordContent        string        `json:"wordContent" bson:"wordContent"  form:"word_content"`
	//	FirstResponseIndex string        `json:"firstResponseIndex" bson:"firstResponseIndex"  form:"fisrtResponseIndex"`
	//	Responses          []models.Comment     `json:"responses" bson:"responses" form:"responses"`
	//}

	if bindJsonErr := ctx.BindJSON(&response); bindJsonErr != nil {
		fmt.Printf("%s\n", bindJsonErr)
	} else {
		response.Id = bson.NewObjectId()
		response.CreatedAt = time.Now()
		response.PublishedAt = time.Now()
		response.ModifiedAt = time.Now()
	}

	query := bson.M{
		"comments.responses._id": bson.ObjectIdHex(ResponseId),
	}

	//pipe := Db.C("picture").Pipe(
	//	[]bson.M{
	//		{
	//			"$match": bson.M{
	//				"comments.responses._id": bson.ObjectIdHex(ResponseId),
	//			},
	//		},
	//		{
	//			"$unwind": "$comments",
	//		},
	//		{
	//			"$group": bson.M{
	//				"_id": "$_id",
	//				"cResponsesIds": bson.M{
	//					"$push": "$comments._id",
	//				},
	//			},
	//		},
	//		{
	//			"$project": bson.M{
	//				"commentId":        "$comments._id",
	//				"wordContent":      "$comments.word_content",
	//				"cResponsesIds":    1,
	//				"commentResponses": "$comments.responses",
	//			},
	//		},
	//		{
	//			"$unwind": "$commentResponses",
	//		},
	//		{
	//			"$project": bson.M{
	//				"commentId":       1,
	//				"wordContent":     "$commentResponses.word_content",
	//				"firstResponseId": "$commentResponses._id",
	//				"firstResponseIndex": bson.M{
	//					"$indexOfArray": []interface{}{
	//						"$cResponsesIds", ResponseId,
	//					},
	//				},
	//				"responses": "$commentResponses.responses",
	//			},
	//		},
	//		{
	//			"$match": bson.M{
	//				"firstResponseId": bson.ObjectIdHex(ResponseId),
	//			},
	//		},
	//
	//	},
	//)
	//pipe := Db.C("picture").Pipe(
	//	// very tricky
	//	[]bson.M{
	//		{
	//			// find corresponding picture
	//			"$match": bson.M{
	//				"comments._id": bson.ObjectIdHex(ResponseId),
	//			},
	//		},
	//		{
	//			"$unwind": "$comments",
	//		},
	//		{
	//			"$project": bson.M{
	//				"commentId":        "$comments._id",
	//				"commentResponses": "$comments.responses",
	//				"wordContent":      "$comments.word_content",
	//			},
	//		},
	//		{
	//			// find the corresponding comments
	//			"$match": bson.M{
	//				"commentId": bson.ObjectIdHex(CommentId),
	//			},
	//		},
	//
	//	},
	//)

	//var result firstResponse
	//
	//if err := pipe.One(&result); err != nil {
	//	fmt.Printf("%s205\n", err)
	//	panic(err)
	//	ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	//	return
	//}
	//fmt.Printf("%v\n", result)
	//
	//tmp_responses := result.Responses
	//fmt.Printf("%s\n", tmp_responses)
	//tmp_responses = append(tmp_responses, response)
	//fmt.Printf("%s\n", tmp_responses)
	//result.Responses = tmp_responses

	setResponse := bson.M{
		"$push": bson.M{
			"comments.$.responses.0.responses": response,
		},
	}

	if picErr := Db.C("picture").Update(query, setResponse); picErr != nil {
		fmt.Printf("%s\n223", picErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	Db.Close()
}

func (this CommentApi) PutCommentToPic(ctx *gin.Context) {
	// TODO: authentication
	id := ctx.Param("id")
	comment := models.Comment{}
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
	}
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

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
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
	} else {
		Db.C("people").Find(bson.M{"comments.id": id}).One(&visitor)
		ctx.JSON(http.StatusOK, visitor.Comments)
	}
}

func (this CommentApi) DeleteComment(ctx *gin.Context) {

}

func (comment *CommentAlias) CommentPreCreateSave(ctx *gin.Context) {
	// TODO: add a global chan to count comment number and other number
	Db := db.MgoDb{}
	Db.Init()

	session := sessions.Default(ctx)
	visitorId := session.Get("visitor").(string)

	comment.Id = bson.NewObjectId()
	if visitorId != "" {
		comment.ById = visitorId
	}
	comment.CreatedAt = time.Now()
	comment.PublishedAt = time.Now()
	comment.ModifiedAt = time.Now()

	Db.Close()
}
