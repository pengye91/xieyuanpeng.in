package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	Id             bson.ObjectId `json:"id" bson:"_id,omitempty"  form:"id"`
	ById           string        `json:"byId" bson:"by_id"  form:"by_id"`
	ByName         string        `json:"byName" bson:"byName"  form:"byName"`
	WordContent    string        `json:"wordContent" bson:"word_content"  form:"word_content"`
	InternalPath   string        `json:"internalPath" bson:"internalPath"  form:"internalPath"`
	ContainPicPath string        `json:"containPicPath" bson:"contain_pic_path"  form:"contain_pic_path"`
	UnderPic       string        `json:"underPic" bson:"under_pic"  form:"under_pic"`
	SlugUrl        string        `json:"slugUrl" bson:"slug_url" form:"slug_url"`
	Comments       []Comment     `json:"comments" bson:"comments" form:"comments"`
	PublishedAt    time.Time     `json:"publishedAt" bson:"published_at" form:"published_at"`
	CreatedAt      time.Time     `json:"createdAt,omitempty" bson:"created_at" form:"created_at"`
	ModifiedAt     time.Time     `json:"modifiedAt,omitempty" bson:"modified_at" form:"updated_at"`
}
