package models

import (
	"time"
)

type Comment struct {
	Id             string    `json:"id" bson:"id"  form:"id"`
	ById           string    `json:"byId" bson:"by_id"  form:"by_id"`
	ByName         string    `json:"byName" bson:"byName"  form:"byName"`
	WordContent    string    `json:"wordContent" bson:"word_content"  form:"word_content"`
	ContainPicPath string    `json:"containPicPath" bson:"contain_pic_path"  form:"contain_pic_path"`
	UnderPic       string    `json:"underPic" bson:"under_pic"  form:"under_pic"`
	SlugUrl        string    `json:"slugUrl" bson:"slug_url" form:"slug_url"`
	Responses      []Comment `json:"responses" bson:"responses" form:"responses"`
	PublishedAt    time.Time `json:"publishedAt" bson:"published_at" form:"published_at"`
	CreatedAt      time.Time `json:"createdAt,omitempty" bson:"created_at" form:"created_at"`
	ModifiedAt     time.Time `json:"modifiedAt,omitempty" bson:"modified_at" form:"updated_at"`
}
