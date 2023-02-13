package service

import (
	"go-gin-course/common/datasource"
	"go-gin-course/pkg/comment/model"
)

type CommentService interface {
	Save(model.Comment) model.Comment
	One(string) (model.Comment, error)
	All() []model.Comment
}

type commentService struct {
}

func New() CommentService {
	return &commentService{}
}

func (service *commentService) One(id string) (model.Comment, error) {
	var comment model.Comment
	err := datasource.DB.First(&comment, id).Error
	return comment, err
}

func (service *commentService) Save(comment model.Comment) model.Comment {
	newItem := model.Comment{
		Content: comment.Content,
	}
	datasource.DB.Create(&newItem)
	return newItem
}

func (service *commentService) All() []model.Comment {
	var items []model.Comment
	datasource.DB.Find(&items)
	return items
}
