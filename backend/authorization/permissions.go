package authorization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
	"github.com/pengye91/xieyuanpeng.in/backend/models"
	"gopkg.in/mgo.v2/bson"
)

var AuthPermissions struct {
	IsAdmin func(string, *gin.Context) bool
}

func IsAdmin(UserId string, ctx *gin.Context) bool {
	Db := &db.MgoDb{}
	Db.Init()
	defer Db.Close()

	user := models.VisitorBasic{}
	fmt.Println(UserId)

	if err := Db.C("auth").FindId(bson.ObjectIdHex(UserId)).One(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Err("5"))
		return false
	} else {
		if user.Name == "xyp" {
			return true
		}
		return false
	}
}

func All(UserId string, ctx *gin.Context) bool {
	return true
}
