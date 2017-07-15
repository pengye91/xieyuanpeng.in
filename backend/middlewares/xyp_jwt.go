package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/mgo.v2/bson"
)

var JWTAuthMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "xyp test",
	Key:        []byte("secret key"),
	Timeout:    time.Hour,
	MaxRefresh: time.Hour,
	Authenticator: func(userID string, password string, ctx *gin.Context) (string, bool) {
		Db := &db.MgoDb{}
		Db.Init()
		defer Db.Close()

		var ps struct {
			Pass string        `json:"pass" bson:"pass" form:"pass"`
		}

		_pass := ctx.PostForm("pass")
		if err := Db.C("auth").FindId(bson.ObjectIdHex(userID)).Select(bson.M{"pass": 1}).One(&ps); err != nil {
			ctx.JSON(http.StatusNotFound, models.Err("1"))
			return userID, false
		}

		pass := libs.Password{}
		cp := pass.Compare(ps.Pass, _pass)

		if cp {
			return userID, true
		} else {
			return userID, false
		}
	},
	Authorizator: func(userID string, c *gin.Context) bool {
		return true
	},
	Unauthorized: func(ctx *gin.Context, code int, message string) {
		ctx.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}
