package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type VisitorNameId map[bson.ObjectId]string

type Picture struct {
	Id          bson.ObjectId   `json:"id" bson:"_id"  form:"id"`
	Title       string          `json:"title" bson:"title"  form:"title"`
	Path        string          `json:"path" bson:"path"  form:"path"`
	Like        int             `json:"like" bson:"like"  form:"like"`
	LikedBy     []VisitorNameId `json:"likedBy" bson:"liked_by"  form:"liked_by"`
	SlugUrl     string          `json:"slug_url" bson:"slug_url" form:"slug_url"`
	Description string          `json:"description" bson:"description"  form:"description"`
	CreatedAt   time.Time       `json:"created_at" bson:"created_at" form:"created_at"`
	Comments    []Comment       `json:"comments" bson:"comments"  form:"comments"`
}
