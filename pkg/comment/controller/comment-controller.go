package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/common"
	"go-gin-course/pkg/comment/model"
	"go-gin-course/pkg/comment/service"
)

type CommentController interface {
	All(ctx *gin.Context)
	Save(ctx *gin.Context)
	One(ctx *gin.Context)
}

type controller struct {
	service service.CommentService
}

func (c controller) All(ctx *gin.Context) {
	models := c.service.All()
	ctx.JSON(200, models)
}

func (c controller) Save(ctx *gin.Context) {
	var comment model.Comment
	err := ctx.BindJSON(&comment)
	if common.ThrowErrorIfExist(ctx, err) {
		return
	}
	ctx.JSON(200, comment)
}

func (c controller) One(ctx *gin.Context) {
	id := ctx.Param("id")
	comment, err := c.service.One(id)
	if common.ThrowErrorIfExist(ctx, err) {
		return
	}
	ctx.JSON(200, comment)
}

func New(service service.CommentService) CommentController {
	return controller{
		service: service,
	}
}
