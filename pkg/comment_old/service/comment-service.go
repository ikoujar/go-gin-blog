package service

import (
	"go-gin-course/pkg/comment_old/model"
)

type CommentService interface {
	Save(model.Comment) model.Comment
	All() []model.Comment
}

type commentService struct {
	comments []model.Comment
}

func (service *commentService) Save(comment model.Comment) model.Comment {
	service.comments = append(service.comments, comment)
	return comment
}

func (service *commentService) All() []model.Comment {
	return service.comments
}

func New() CommentService {
	return &commentService{}
}
