package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/mgo.v2/bson"
)

type AuthAPI struct {
	*gin.Context
}

type LoginInfo struct {
	LoginId string `json:"logId"`
	Pass    string `json:"pass"`
}

func (this AuthAPI) Register(ctx *gin.Context) {
	session := sessions.Default(ctx)

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	visitorInfo := models.VisitorBasic{}
	err := ctx.BindJSON(&visitorInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Err("4"))
		return
	}

	pass := libs.Password{}
	visitorInfo.Pass = pass.Gen(visitorInfo.Pass)

	visitorInfo.Id = bson.NewObjectId()
	visitorInfo.CreatedAt = time.Now()
	visitorInfo.UpdatedAt = time.Now()

	// Insert Visitor
	if err := Db.C("auth").Insert(&visitorInfo); err != nil {
		// Is a duplicate key, but we don't know which one
		ctx.JSON(http.StatusOK, models.Err("5"))
		if Db.IsDup(err) {
			ctx.JSON(http.StatusOK, models.Err("6"))
		}
	} else {
		visitor := models.Visitor{}
		visitor.Basic = visitorInfo
		visitor.Id = visitorInfo.Id
		insertToPeopleErr := Db.C("people").Insert(&visitor)
		if insertToPeopleErr != nil {
			ctx.JSON(http.StatusBadRequest, models.Err("5"))
			return
		}
		// auto login
		session.Set("login", "true")
		session.Set("visitor", visitorInfo.Id.String())
		session.Save()

		ctx.JSON(http.StatusOK, visitorInfo)
	}

}

func (this AuthAPI) Login(ctx *gin.Context) {
	session := sessions.Default(ctx)

	result := models.VisitorBasic{}
	loginInfo := LoginInfo{}
	err := ctx.BindJSON(&loginInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
		return
	}
	_loginId := string(loginInfo.LoginId)
	_pass := string(loginInfo.Pass)

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if strings.Contains(_loginId, "@") {
		if err := Db.C("auth").Find(bson.M{"email": _loginId}).One(&result); err != nil {
			ctx.JSON(http.StatusNotFound, models.Err("1"))
			return
		}
	} else {
		if err := Db.C("auth").Find(bson.M{"name": _loginId}).One(&result); err != nil {
			ctx.JSON(http.StatusNotFound, models.Err("1"))
			return
		}
	}

	pass := libs.Password{}
	var cp = pass.Compare(result.Pass, _pass)

	if cp {
		token := pass.Token()
		session.Set("login", "true")
		session.Set("visitor", result.Id.String())
		session.Save()
		ctx.JSON(http.StatusOK, gin.H{"response": true, "token": token})
	} else {
		ctx.JSON(http.StatusOK, models.Err("7"))
	}
}

func (this AuthAPI) Check(ctx *gin.Context) {
	session := sessions.Default(ctx)
	Db := &db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var ps struct {
		Pass string        `json:"pass" bson:"pass" form:"pass"`
	}

	_pass := string(ctx.PostForm("pass"))
	if visitor := session.Get("visitor"); visitor == nil {
		ctx.JSON(http.StatusInternalServerError, "session error")
		return
	} else if err := Db.C("auth").FindId(bson.ObjectIdHex(visitor.(string))).Select(bson.M{"pass": 1}).One(&ps); err != nil {
		ctx.JSON(http.StatusNotFound, models.Err("1"))
		return
	}

	pass := libs.Password{}
	cp := pass.Compare(ps.Pass, _pass)

	if cp {
		ctx.JSON(http.StatusOK, gin.H{"response": true})
	} else {
		ctx.JSON(http.StatusOK, models.Err("8"))
	}

}

func (this AuthAPI) GetAllUsers(ctx *gin.Context) {
	visitors := []models.VisitorBasic{}

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("auth").Find(nil).All(&visitors); err != nil {
		ctx.JSON(http.StatusNotFound, models.Err("1"))
		return
	}
	ctx.JSON(http.StatusOK, visitors)

}
