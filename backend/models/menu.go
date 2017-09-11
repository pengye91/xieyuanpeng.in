package models

import (
	"gopkg.in/mgo.v2/bson"
)

type SideMenuItems map[string]map[string]string
type AdminSideMenuItems map[string]map[string]string

type MenuItem struct {
	Ref                string            `json:"ref" bson:"ref" form:"ref"`
	Name               string            `json:"name" bson:"name" form:"name"`
	SideMenuItems      map[string]string `json:"sideMenuItems" bson:"sideMenuItems" form:"side_menu_items"`
	AdminSideMenuItems map[string]string `json:"adminSideMenuItems" bson:"adminSideMenuItems" form:"admin_side_menu_items"`
}

type Menu struct {
	Id        bson.ObjectId       `json:"id" bson:"_id,omitempty" form:"id"`
	MenuItems map[string]MenuItem `json:"menuItems" bson:"menuItems" form:"menu_items"`
}
