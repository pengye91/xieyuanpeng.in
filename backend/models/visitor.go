package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// VisitorBasic represents the visitors' basic info of foto.in
	VisitorBasic struct {
		Id        bson.ObjectId `json:"id" bson:"_id"  form:"id"`
		Name      string        `json:"name" bson:"name"  form:"name"`
		Email     string        `json:"email" bson:"email"  form:"email"`
		Pass      string        `json:"pass" bson:"pass" form:"pass"`
		CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at" form:"created_at"`
		UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at" form:"updated_at"`
	}

	VisitorNameId map[bson.ObjectId]string

	Visitor struct {
		Id       bson.ObjectId `json:"id" bson:"_id"  form:"id"`
		Basic    VisitorBasic  `json:"basic" bson:"basic" form:"basic"`
		Messages []Message     `json:"messages" bson:"messages"  form:"messages"`
		Comments []Comment     `json:"comments" bson:"comments"  form:"comments"`
	}

	Message struct {
		Id        bson.ObjectId `json:"id" bson:"_id"  form:"id"`
		SlugUrl   string        `json:"slug_url" bson:"slug_url" form:"slug_url"`
		Content   string        `json:"content" bson:"content"  form:"content"`
		CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at" form:"created_at"`
		UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at" form:"updated_at"`
	}
)
