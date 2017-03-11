package api

import (
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"

	"github.com/pengye91/xieyuanpeng.in/mongo/backend/db"
	"github.com/pengye91/xieyuanpeng.in/mongo/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/mongo/backend/models"
)

type AuthAPI struct {
	*iris.Context
}

type LoginInfo struct {
	Email  string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
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

	// Insert Visitor
	if err := Db.C("auth").Insert(&visitorInfo); err != nil {
		// Is a duplicate key, but we don't know which one
		ctx.JSON(iris.StatusOK, models.Err("5"))
		if Db.IsDup(err) {
			ctx.JSON(iris.StatusOK, models.Err("6"))
		}
	} else {
		ctx.JSON(iris.StatusOK, iris.Map{"response": true})
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
	_email := string(loginInfo.Email)
	_pass := string(loginInfo.Password)

	Db := db.MgoDb{}
	Db.Init()

	if err := Db.C("auth").Find(bson.M{"email": _email}).One(&result); err != nil {
		ctx.JSON(iris.StatusOK, models.Err("1"))
		return
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
