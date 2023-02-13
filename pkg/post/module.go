package post

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/common/datasource"
	"go-gin-course/pkg/post/controller"
	"go-gin-course/pkg/post/model"
	"go-gin-course/pkg/post/service"
)

//var (
//	//PostService    = service.New()
//	PostController = controller.New(service.New())
//)

func Init(route *gin.Engine) {

	// Run datasource migration.
	datasource.AutoMigrate(&model.Post{})

	// Initialize controller.

	PostController := controller.New(service.New())

	// Creat router group.
	router := route.Group("/posts")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, PostController.All())
	})
	router.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, PostController.Save(ctx))
	})

	router.GET("/:id", func(ctx *gin.Context) {
		PostController.One(ctx)
	})

}
