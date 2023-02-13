package comment

import (
	"github.com/gin-gonic/gin"
	"go-gin-course/common/datasource"
	"go-gin-course/pkg/comment/controller"
	"go-gin-course/pkg/comment/model"
	"go-gin-course/pkg/comment/service"
)

//var (
//	//CommentService    = service.New()
//	CommentController = controller.New(service.New())
//)

func Init(route *gin.Engine) {

	// Run datasource migration.
	datasource.AutoMigrate(&model.Comment{})

	// Initialize controller.
	CommentController := controller.New(service.New())

	// Creat router group.
	router := route.Group("/comments")

	router.GET("/", CommentController.All)
	router.POST("/", CommentController.Save)
	router.GET("/:id", CommentController.One)

}
