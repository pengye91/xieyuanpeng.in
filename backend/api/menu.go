package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pengye91/xieyuanpeng.in/backend/db"
)

// Todo: Add MenuItems API. This should the root of all front-end classification.

type MenuApi struct {
	*gin.Context
}

func (this MenuApi) GetMenuItems(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


}

func (this MenuApi) PutMenuItems(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

}

func (this MenuApi) PostMenuItems(ctx *gin.Context) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

}
