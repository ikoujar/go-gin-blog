package model

import "go-gin-course/pkg/post/model"

type Comment struct {
	ID      string     `json:"id" gorm:"primary_key"`
	PostID  int        `json:"post_id"`
	Post    model.Post `gorm:"references:ID"` // use Post.PostID as references
	Content string     `json:"content"`
}
