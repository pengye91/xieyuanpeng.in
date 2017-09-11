package middlewares

import (
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"github.com/pengye91/xieyuanpeng.in/backend/utils/log"
	"gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/mgo.v2/bson"
)

var user models.VisitorBasic

type passUser struct {
	Id    bson.ObjectId `json:"id" bson:"_id"  form:"id"`
	Name  string        `json:"name" bson:"name"  form:"name"`
	Email string        `json:"email" bson:"email"  form:"email"`
}

const (
	Month    = 30 * 24 * time.Hour
	Year     = 30 * 24 * time.Hour * 12
	TenYears = 30 * 24 * time.Hour * 12 * 10
)

var JWTAuthMiddleware = jwt.GinJWTMiddleware{
	Realm:      "xyp",
	Key:        []byte("secret key"),
	Timeout:    TenYears,
	MaxRefresh: TenYears,

	// for loginHandler usage
	Authenticator: func(loginID string, password string, ctx *gin.Context) (string, bool) {
		Db := &db.MgoDb{}
		Db.Init()
		defer Db.Close()
		session := sessions.Default(ctx)

		if strings.Contains(loginID, "@") {
			if err := Db.C("auth").Find(bson.M{"email": loginID}).One(&user); err != nil {
				log.LoggerSugar.Errorw("JWTAuthMiddleware email Login Error",
					"module", "jwt",
					"error", err,
				)
				return loginID, false
			}
		} else {
			if err := Db.C("auth").Find(bson.M{"name": loginID}).One(&user); err != nil {
				log.LoggerSugar.Errorw("JWTAuthMiddleware name Login Error",
					"module", "application: jwt",
					"error", err,
				)
				return loginID, false
			}
		}

		pass := libs.Password{}
		cp := pass.Compare(user.Pass, password)

		if cp {
			session.Set("logined", "true")
			session.Set("visitor", user.Id.String())
			session.Save()
			return user.Id.Hex(), true
		} else {
			return user.Id.Hex(), false
		}
	},

	// On every JWT related handler
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
		return map[string]interface{}{"user": passUser{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}}
	},
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}

func JWTMiddlewareFactory(authorizator func(string, *gin.Context) bool) *jwt.GinJWTMiddleware {
	authorizatorVaryMiddleware := JWTAuthMiddleware
	authorizatorVaryMiddleware.Authorizator = authorizator
	return &authorizatorVaryMiddleware
}
