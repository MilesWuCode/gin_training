package router

import (
	"github.com/gin-gonic/gin"

	"mygo/controller"
)

func Init() {
	router := router()

	router.Run()
}

func router() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		userController := controller.UserController{}

		users.GET("", userController.Index)
		users.GET("/:id", userController.Show)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	books := router.Group("/books")
	{
		bookController := controller.BookController{}

		books.GET("", bookController.Index)
		books.GET("/:id", bookController.Show)
		books.POST("", bookController.Create)
		books.PUT("/:id", bookController.Update)
		books.DELETE("/:id", bookController.Delete)
	}

	return router
}
