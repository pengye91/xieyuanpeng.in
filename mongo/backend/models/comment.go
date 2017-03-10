package models

import (
	"time"
)

type Comment struct {
	Id          string       `json:"string,id" bson:"_id"  form:"id"`
	By          VisitorBasic `json:"string,by" bson:"by"  form:"by"`
	WordContent string       `json:"string,word_content" bson:"word_content"  form:"word_content"`
	PicPath     string       `json:"string,pic_path" bson:"pic_path"  form:"pic_path"`
	SlugUrl     string       `json:"slug_url, string" bson:"slug_url" form:"slug_url"`
	Responses   []Comment    `json:"responses, string" bson:"responses" form:"responses"`
	PublishedAt time.Time    `json:"published_at, string" bson:"published_at" form:"published_at"`
	CreatedAt   time.Time    `json:"created_at,omitempty, string" bson:"created_at" form:"created_at"`
	ModifiedAt  time.Time    `json:"modified_at,omitempty, string" bson:"modified_at" form:"updated_at"`
}
