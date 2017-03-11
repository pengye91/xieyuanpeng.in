package api

import (
	"gopkg.in/kataras/iris.v5"
	"gopkg.in/mgo.v2/bson"

	"github.com/pengye91/xieyuanpeng.in/mongo/backend/models"
	"github.com/pengye91/xieyuanpeng.in/mongo/backend/db"
)

type UserAPI struct {
	*iris.Context
}

// GET /users
func (this UserAPI) Get() {
	Db := db.MgoDb{}
	Db.Init()

	visitors := []models.Visitor{}
	if err := Db.C("people").Find(nil).All(&visitors); err != nil {
		this.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		this.JSON(iris.StatusOK, &visitors)
	}

	Db.Close()
}

// GET /users/:param1
func (this UserAPI) GetBy(id string) {
	Db := db.MgoDb{}
	Db.Init()
	visitor := models.Visitor{}

	if err := Db.C("people").Find(bson.M{"id": id}).One(&visitor); err != nil {
		this.JSON(iris.StatusOK, models.Err("1"))
		return
	} else {
		this.JSON(iris.StatusOK, &visitor)
	}

	Db.Close()

}

// PUT /users/:param1
func (this UserAPI) PutBy(id string) {

	Db := db.MgoDb{}
	Db.Init()

	//	result := models.User{}

	name := string(this.FormValue("name"))
	email := string(this.FormValue("email"))

	//	if err := Db.C("people").Update(bson.M{"id": id}, bson.M{"$set": bson.M{"name": name}}); err != nil {
	//		this.JSON(iris.StatusOK, models.Err("5"))
	//	} else {
	//		println(name + " has been inserted to database")
	//		this.JSON(iris.StatusOK, iris.Map{"response": true})

	//	}
	//	Db.InsertUser(name)

	// Update
	c := Db.C("people")
	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": bson.M{"name": name, "email": email}}
	err := c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	} else {
		println(name + " has been inserted to database")
		this.JSON(iris.StatusOK, iris.Map{"response": true})

	}

}

// POST /users/:param1
func (this UserAPI) PostBy(id string) {

	usr := models.User{}
	err := this.ReadForm(&usr)

	if err != nil {
		this.JSON(iris.StatusOK, models.Err("4"))
		panic(err.Error())
	}

	usr.Id = id

	Db := db.MgoDb{}
	Db.Init()

	// Insert
	if err := Db.C("people").Insert(&usr); err != nil {
		this.JSON(iris.StatusOK, models.Err("5"))
	} else {
		this.JSON(iris.StatusOK, iris.Map{"response": true})
	}

	Db.Close()

}

// DELETE /users/:param1
func (this UserAPI) DeleteBy(id string) {
	Db := db.MgoDb{}
	Db.Init()

	if err := Db.C("people").Remove(bson.M{"id": id}); err != nil {
		this.JSON(iris.StatusOK, models.Err("1"))
	} else {

		this.JSON(iris.StatusOK, iris.Map{"response": true})
	}

}

// // Get Params example code
// var _name = string(this.FormValue("name"))
// var _grender = string(this.FormValue("gender"))
// var _age = string(this.FormValue("age"))
// _newage, _ := strconv.Atoi(_age)
// var _id = string(this.FormValue("id"))

// usr := models.User{
// 	Name:   _name,
// 	Email:   "ivancduran@gmail.com",
// 	Gender: false,
// 	Birth:  "1989-08-21",
// 	Id:     _id,
// }
