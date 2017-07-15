package middlewares

import (
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/mgo.v2/bson"
)
var user models.VisitorBasic

var JWTAuthMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "xyp test",
	Key:        []byte("secret key"),
	Timeout:    time.Hour,
	MaxRefresh: time.Hour,
	Authenticator: func(loginID string, password string, ctx *gin.Context) (string, bool) {
		Db := &db.MgoDb{}
		Db.Init()
		defer Db.Close()

		session := sessions.Default(ctx)

		if strings.Contains(loginID, "@") {
			if err := Db.C("auth").Find(bson.M{"email": loginID}).One(&user); err != nil {
				return loginID, false
			}
		} else {
			if err := Db.C("auth").Find(bson.M{"name": loginID}).One(&user); err != nil {
				return loginID, false
			}
		}

		pass := libs.Password{}
		cp := pass.Compare(user.Pass, password)

		if cp {
			session.Set("logined", "true")
			session.Set("visitor", user.Id.String())
			session.Save()
			return user.Id.String(), true
		} else {
			return user.Id.String(), false
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
	PayloadFunc: func(userID string) map[string]interface{} {
		return map[string]interface{}{"user": user}
	},
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}
