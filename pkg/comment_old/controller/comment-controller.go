package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/pkg/comment_old/model"
	"go-gin-course/pkg/comment_old/service"
)

type CommentController interface {
	All() []model.Comment
	Save(ctx *gin.Context) model.Comment
}

type controller struct {
	service service.CommentService
}

func (c controller) All() []model.Comment {
	return c.service.All()
}

func (c controller) Save(ctx *gin.Context) model.Comment {
	var comment model.Comment
	ctx.BindJSON(&comment)
	return c.service.Save(comment)
}

func New(service service.CommentService) CommentController {
	return controller{
		service: service,
	}
}
