package api

import (
	"net/http"
	"strings"
	"time"

<<<<<<< HEAD
=======
	"fmt"
>>>>>>> dev
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/libs"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/mgo.v2/bson"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> dev
)

var tk string

type AuthAPI struct {
	*gin.Context
}
type Set map[interface{}]bool

var (
	UsernameSet = make(Set)
	EmailSet    = make(Set)
)

type LoginInfo struct {
	LoginId string `json:"logId"`
	Pass    string `json:"pass"`
}

func (this AuthAPI) Register(ctx *gin.Context) {
<<<<<<< HEAD
	visitorInfo := models.VisitorBasic{}
	err := ctx.BindJSON(&visitorInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Err("4"))
=======
	session := sessions.Default(ctx)

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	visitor := models.VisitorBasic{}
	if err := ctx.BindJSON(&visitor); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err("4"))
>>>>>>> dev
		return
	}

	pass := libs.Password{}
	visitor.Pass = pass.Gen(visitor.Pass)

	visitor.Id = bson.NewObjectId()
	visitor.CreatedAt = time.Now()
	visitor.UpdatedAt = time.Now()

	// Insert Visitor
	if err := Db.C("auth").Insert(&visitor); err != nil {
		// Is a duplicate key, but we don't know which one
<<<<<<< HEAD

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
=======
		ctx.JSON(http.StatusBadRequest, models.Err("5"))
		return
	} else {
		// TODO: auto login after registration.
		session.Set("login", "true")
		session.Set("visitor", visitor.Id.String())
		session.Save()

		UsernameSet[visitor.Name] = true
		EmailSet[visitor.Email] = true

		ctx.JSON(http.StatusCreated, visitor)
>>>>>>> dev
	}
}

func (this AuthAPI) Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
<<<<<<< HEAD

=======
>>>>>>> dev
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
<<<<<<< HEAD
		session.Set("visitor", result.Id.Hex())
		session.Set("token", token)

		if err := session.Save(); err != nil {
			fmt.Printf("%s", err)
		}

=======
		session.Set("visitor", result)
		session.Save()
		//time.Sleep(3 * time.Second)
>>>>>>> dev
		ctx.JSON(http.StatusOK, gin.H{"response": true, "token": token})
	} else {
		ctx.JSON(http.StatusBadRequest, models.Err("7"))
	}
}

func (this AuthAPI) Check(ctx *gin.Context) {
<<<<<<< HEAD
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
=======
	session := sessions.Default(ctx)
	Db := &db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var ps struct {
		Pass string `json:"pass" bson:"pass" form:"pass"`
	}

	_pass := string(ctx.PostForm("pass"))
	visitor := session.Get("visitor")
	if visitor == nil {
		ctx.JSON(http.StatusInternalServerError, "session error")
		return
	} else {
		fmt.Println(visitor)
		if err := Db.C("auth").FindId(visitor).Select(bson.M{"pass": 1}).One(&ps); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusNotFound, models.Err("1"))
			return
		}
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
>>>>>>> dev
	}
	ctx.JSON(http.StatusOK, visitors)

}

<<<<<<< HEAD
func (this AuthAPI) Session(ctx *gin.Context) {
	session := sessions.Default(ctx)

	login := session.Get("login").(string)
	token := session.Get("token").(string)

	if login == "true" {
		ctx.JSON(http.StatusOK, gin.H{"response": true, "token": token})
	} else {
		ctx.JSON(http.StatusOK, models.Err("8"))
=======
func (this AuthAPI) LogOut(ctx *gin.Context) {
	for _, v := range ctx.Request.Cookies() {
		fmt.Println(v)
	}
	ctx.SetCookie("sessionid", "", -1, "/", "localhost", false, false)
	ctx.JSON(http.StatusOK, gin.H{"OK": "DONE"})
}

func AutoSearch(ctx *gin.Context) {
	username := ctx.Query("username")
	email := ctx.Query("email")

	if username != "" {
		if UsernameSet[username] {
			ctx.JSON(http.StatusOK, gin.H{
				username: "Registered",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{
				username: "Not Registered",
			})
		}
	} else if email != "" {
		if EmailSet[email] {
			ctx.JSON(http.StatusOK, gin.H{
				email: "Ooops, Registered",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{
				email: "OK, Not Registered",
			})

		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"WRONG QUERY": "Query params must contain one of username and email.",
		})
>>>>>>> dev
	}
}

// TODO: Put the logic in redis. Or a small search DB.
func InitialSetsFromDB() {
	Db := &db.MgoDb{}
	Db.Init()
	defer Db.Close()
	var (
		usernames []struct {
			Name string `json:"name" bson:"name"  form:"name"`
		}
		emails []struct {
			Email string `json:"email" bson:"email"  form:"email"`
		}
	)

	if usernameErr := Db.C("auth").Find(nil).Select(bson.M{"name": 1}).All(&usernames); usernameErr != nil {
		fmt.Println(usernameErr)
	}
	if emailErr := Db.C("auth").Find(nil).Select(bson.M{"email": 1}).All(&emails); emailErr != nil {
		fmt.Println(emailErr)
	}

	for _, v := range usernames {
		UsernameSet[v.Name] = true
	}

	for _, v := range emails {
		EmailSet[v.Email] = true
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
