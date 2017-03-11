package models

import (
	"time"
)

type Picture struct {
	Id          string           `json:"string,id" bson:"_id"  form:"id"`
	Title       string           `json:"string,title" bson:"title"  form:"title"`
	Path        string           `json:"string,path" bson:"path"  form:"path"`
	Like        int              `json:"string,like" bson:"like"  form:"like"`
	SlugUrl     string           `json:"slug_url, string" bson:"slug_url" form:"slug_url"`
	Description string           `json:"string,description" bson:"description"  form:"description"`
	CreatedAt   time.Time        `json:"created_at,omitempty, string" bson:"created_at" form:"created_at"`
	Comments    []models.Comment `json:"string,comments" bson:"comments"  form:"comments"`
}
