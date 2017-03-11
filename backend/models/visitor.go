package models

import (
	"time"
)

type (
	// VisitorBasic represents the visitors' basic info of foto.in
	VisitorBasic struct {
		Id       string       `json:"id" bson:"id"  form:"id"`
		Name      string    `json:"name" bson:"name"  form:"name"`
		Email     string    `json:"email" bson:"email"  form:"email"`
		Pass      string    `json:"pass" bson:"pass" form:"pass"`
		CreatedAt time.Time `json:"created_at,omitempty, string" bson:"created_at" form:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty, string" bson:"updated_at" form:"updated_at"`
	}

	Visitor struct {
		Id       string       `json:"id" bson:"id"  form:"id"`
		Basic    VisitorBasic `json:"basic" bson:"basic" form:"basic"`
		Messages []Message    `json:"messages" bson:"messages"  form:"messages"`
		Comments []Comment    `json:"comments" bson:"comments"  form:"comments"`
	}

	Message struct {
		Id        string    `json:"id" bson:"id"  form:"id"`
		SlugUrl   string    `json:"slug_url" bson:"slug_url" form:"slug_url"`
		Content   string    `json:"content" bson:"content"  form:"content"`
		CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at" form:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at" form:"updated_at"`
	}
)
