package comment_old

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/pkg/comment_old/controller"
	"go-gin-course/pkg/comment_old/service"
)

var (
	//PostService    = service.New()
	CommentController = controller.New(service.New())
)

func Init(route *gin.Engine) {
	router := route.Group("/comments")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, CommentController.All())
	})
	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, CommentController.Save(ctx))
	})
}
