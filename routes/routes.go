package routes

import (
	"localSandbox/SimpleCRUD/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	userController := controller.NewUserController()
	postController := controller.NewPostController()
	commentController := controller.NewCommentController()

	r := gin.Default()
	r.GET("/users/getAll", userController.GetAll)
	r.GET("/users/get/:id", userController.Get)
	r.POST("/users/add", userController.Add)
	r.POST("/users/update", userController.Update)
	r.DELETE("/users/delete/:id", userController.Delete)
	r.GET("/posts/getAll", postController.GetAll)
	r.GET("/posts/get/:id", postController.Get)
	r.POST("/posts/add", postController.Add)
	r.GET("/comments/getAll", commentController.GetAll)
	r.GET("/comments/get/:id", commentController.Get)
	r.POST("/comments/add", commentController.Add)
	r.Run()
}
