package models

import (
	"time"
)

type Picture struct {
	Id          string    `json:"id" bson:"id"  form:"id"`
	Title       string    `json:"title" bson:"title"  form:"title"`
	Path        string    `json:"path" bson:"path"  form:"path"`
	Like        int       `json:"like" bson:"like"  form:"like"`
	SlugUrl     string    `json:"slug_url" bson:"slug_url" form:"slug_url"`
	Description string    `json:"description" bson:"description"  form:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at" form:"created_at"`
	Comments    []Comment `json:"comments" bson:"comments"  form:"comments"`
}
