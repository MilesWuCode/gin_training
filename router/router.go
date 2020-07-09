package router

import (
	"github.com/gin-gonic/gin"

	"mygo/controller"
)

func Init() {
	router := Router()

	router.Run()
}

func Router() *gin.Engine {
	router := gin.Default()

	userController := controller.UserController{}

	users := router.Group("/users")
	{
		users.GET("", userController.Index)
		users.GET("/:id", userController.Show)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	bookController := controller.BookController{}

	books := router.Group("/books")
	{
		books.GET("", bookController.Index)
		books.GET("/:id", bookController.Show)
		books.POST("", bookController.Create)
		books.PUT("/:id", bookController.Update)
		books.DELETE("/:id", bookController.Delete)
	}

	return router
}
