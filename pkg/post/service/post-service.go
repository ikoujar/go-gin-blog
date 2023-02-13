package service

import (
	"go-gin-course/common/datasource"
	"go-gin-course/pkg/post/model"
)

type PostService interface {
	Save(model.Post) model.Post
	One(string) (model.Post, error)
	All() []model.Post
}

type postService struct {
}

func New() PostService {
	return &postService{}
}

func (service *postService) One(id string) (model.Post, error) {
	var post model.Post
	err := datasource.DB.First(&post, id).Error
	return post, err
}

func (service *postService) Save(post model.Post) model.Post {
	newItem := model.Post{
		Title:       post.Title,
		Description: post.Description,
		Url:         post.Url,
	}
	datasource.DB.Create(&newItem)
	return newItem
}

func (service *postService) All() []model.Post {
	var items []model.Post
	datasource.DB.Find(&items)
	return items
}
