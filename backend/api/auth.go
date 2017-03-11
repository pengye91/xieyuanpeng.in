package api

import (
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"

	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"strconv"
	"strings"
	"time"
)

type AuthAPI struct {
	*iris.Context
}

type LoginInfo struct {
	LoginId string `json:"logId"`
	Pass    string `json:"pass"`
}

func (this AuthAPI) Register(ctx *iris.Context) {
	visitorInfo := models.VisitorBasic{}
	err := ctx.ReadJSON(&visitorInfo)
	if err != nil {
		ctx.JSON(iris.StatusOK, models.Err("4"))
		return
	}

	pass := libs.Password{}
	visitorInfo.Pass = pass.Gen(visitorInfo.Pass)

	Db := db.MgoDb{}
	Db.Init()

	visitorNumber, _ := Db.C("auth").Count()
	visitorInfo.Id = strconv.Itoa(visitorNumber + 1)
	visitorInfo.CreatedAt = time.Now()
	visitorInfo.UpdatedAt = time.Now()

	// Insert Visitor
	if err := Db.C("auth").Insert(&visitorInfo); err != nil {
		// Is a duplicate key, but we don't know which one
		ctx.JSON(iris.StatusOK, models.Err("5"))
		if Db.IsDup(err) {
			ctx.JSON(iris.StatusOK, models.Err("6"))
		}
	} else {
		visitor := models.Visitor{}
		visitor.Basic = visitorInfo
		visitor.Id = visitorInfo.Id
		insertToPeopleErr := Db.C("people").Insert(&visitor)
		if insertToPeopleErr != nil {
			ctx.JSON(iris.StatusBadRequest, models.Err("5"))
			return
		}
		ctx.JSON(iris.StatusOK, visitorInfo)
	}
	Db.Close()
}

func (this AuthAPI) Login(ctx *iris.Context) {
	result := models.VisitorBasic{}
	loginInfo := LoginInfo{}
	err := ctx.ReadJSON(&loginInfo)
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, models.Err("5"))
		return
	}
	_loginId := string(loginInfo.LoginId)
	_pass := string(loginInfo.Pass)

	Db := db.MgoDb{}
	Db.Init()

	if strings.Contains(_loginId, "@") {
		if err := Db.C("auth").Find(bson.M{"email": _loginId}).One(&result); err != nil {
			ctx.JSON(iris.StatusNotFound, models.Err("1"))
			return
		}
	} else {
		if err := Db.C("auth").Find(bson.M{"name": _loginId}).One(&result); err != nil {
			ctx.JSON(iris.StatusNotFound, models.Err("1"))
			return
		}
	}

	pass := libs.Password{}
	var cp = pass.Compare(result.Pass, _pass)

	if cp {
		token := pass.Token()
		ctx.Session().Set("login", "true")
		ctx.Session().Set("token", token)
		ctx.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
	} else {
		ctx.JSON(iris.StatusOK, models.Err("7"))
	}
	Db.Close()
}

func (this AuthAPI) Check(ctx *iris.Context) {

	_pass := string(ctx.FormValue("pass"))
	token := ctx.Session().GetString("token")

	pass := libs.Password{}
	cp := pass.Compare(token, _pass)

	if cp {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
	} else {
		ctx.JSON(iris.StatusOK, models.Err("8"))
	}

}

func (this AuthAPI) Session(ctx *iris.Context) {

	login := ctx.Session().GetString("login")
	token := ctx.Session().GetString("token")

	if login == "true" {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true, "token": token})
	} else {
		ctx.JSON(iris.StatusOK, models.Err("8"))
	}

}
