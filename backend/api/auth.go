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
	"fmt"
)

var tk string

type AuthAPI struct {
	*gin.Context
}

type LoginInfo struct {
	LoginId string `json:"logId"`
	Pass    string `json:"pass"`
}

func (this AuthAPI) Register(ctx *gin.Context) {
	visitorInfo := models.VisitorBasic{}
	err := ctx.BindJSON(&visitorInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Err("4"))
		return
	}

	pass := libs.Password{}
	visitorInfo.Pass = pass.Gen(visitorInfo.Pass)

	Db := db.MgoDb{}
	Db.Init()

	visitorInfo.Id = bson.NewObjectId()
	visitorInfo.CreatedAt = time.Now()
	visitorInfo.UpdatedAt = time.Now()

	// Insert Visitor
	if err := Db.C("auth").Insert(&visitorInfo); err != nil {
		// Is a duplicate key, but we don't know which one

		ctx.JSON(http.StatusBadRequest, err)
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
		ctx.JSON(http.StatusOK, visitorInfo)
	}
	Db.Close()
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
		session.Set("visitor", result.Id.Hex())
		session.Set("token", token)

		if err := session.Save(); err != nil {
			fmt.Printf("%s", err)
		}

		ctx.JSON(http.StatusOK, gin.H{"response": true, "token": token})
	} else {
		ctx.JSON(http.StatusBadRequest, models.Err("7"))
	}
}

func (this AuthAPI) Check(ctx *gin.Context) {
	start := time.Now()
	session := sessions.Default(ctx)
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	//var p struct {
	//	Pass string        `json:"pass" bson:"pass" form:"pass"`
	//}

	v := models.VisitorBasic{}

	_pass := ctx.PostForm("pass")
	token := session.Get("token").(string)
	visitorId := session.Get("visitor").(string)
	fmt.Printf("%s", visitorId)

	if err := Db.C("auth").FindId(bson.ObjectIdHex(visitorId)).One(&v); err != nil {
		panic(err)
	}


	//if err := Db.C("auth").FindId(bson.ObjectIdHex(visitorId)).Select(bson.M{"pass": 1}).One(&p); err != nil {
	//	panic(err)
	//}

	passLib := libs.Password{}

	cp := passLib.Compare(v.Pass, _pass)
	fmt.Println(v.Pass)

	if cp {
		fmt.Println(time.Since(start))
		ctx.JSON(http.StatusOK, gin.H{"response": true, "token": token})
	} else {
		ctx.JSON(http.StatusOK, models.Err("8"))
	}

}

func (this AuthAPI) Session(ctx *gin.Context) {
	session := sessions.Default(ctx)

	login := session.Get("login").(string)
	token := session.Get("token").(string)

	if login == "true" {
		ctx.JSON(http.StatusOK, gin.H{"response": true, "token": token})
	} else {
		ctx.JSON(http.StatusOK, models.Err("8"))
	}

}
func Xixihaha(ctx *gin.Context) {

	var userId string
	session := sessions.Default(ctx)
	user := session.Get("userID")
	if user == nil {
		fmt.Println("user is nil")
	} else {
		userId = user.(string)
	}
	session.Set("userID", "1234567890")
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{"userId": userId})

}
