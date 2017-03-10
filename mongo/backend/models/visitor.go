package models

import (
	"time"
)

type (
	// VisitorBasic represents the visitors' basic info of foto.in
	VisitorBasic struct {
		Id        string    `json:"string,id" bson:"_id"  form:"id"`
		Name      string    `json:"string,name" bson:"name"  form:"name"`
		Email     string    `json:"string,email" bson:"email"  form:"email"`
		Pass      string    `json:"string,pass" bson:"pass" form:"pass"`
		CreatedAt time.Time `json:"created_at,omitempty, string" bson:"created_at" form:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty, string" bson:"updated_at" form:"updated_at"`
	}

	Visitor struct {
		Basic    VisitorBasic     `json:"string, basic" bson:"basic" form:"basic"`
		Messages []models.Message `json:"string,messages" bson:"messages"  form:"messages"`
	}

	Message struct {
		SlugUrl   string    `json:"slug_url, string" bson:"slug_url" form:"slug_url"`
		Content   string    `json:"string,content" bson:"content"  form:"content"`
		CreatedAt time.Time `json:"created_at,omitempty, string" bson:"created_at" form:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty, string" bson:"updated_at" form:"updated_at"`
	}
)
