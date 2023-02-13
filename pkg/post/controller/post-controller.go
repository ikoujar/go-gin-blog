package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/common"
	"go-gin-course/pkg/post/model"
	"go-gin-course/pkg/post/service"
)

type PostController interface {
	All() []model.Post
	Save(ctx *gin.Context) model.Post
	One(ctx *gin.Context)
}

type controller struct {
	service service.PostService
}

func (c controller) All() []model.Post {
	return c.service.All()
}

func (c controller) Save(ctx *gin.Context) model.Post {
	var post model.Post
	err := ctx.BindJSON(&post)
	if err != nil {
		return model.Post{}
	}
	return c.service.Save(post)
}

func (c controller) One(ctx *gin.Context) {
	id := ctx.Param("id")
	post, err := c.service.One(id)
	if common.ThrowErrorIfExist(ctx, err) {
		return
	}
	ctx.JSON(200, post)
}

func New(service service.PostService) PostController {
	return controller{
		service: service,
	}
}
