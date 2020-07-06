package controller

import (
	"mygo/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BookController struct{}

func (controller BookController) Index(c *gin.Context) {
	var bookService service.BookService

	bookData, err := bookService.GetAll()

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(404, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, bookData)
	}
}

func (controller BookController) Show(c *gin.Context) {
	id := c.Params.ByName("id")

	var bookService service.BookService

	bookData, err := bookService.Get(id)

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(404, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, bookData)
	}
}

func (controller BookController) Create(c *gin.Context) {
	var bookService service.BookService

	bookData, err := bookService.Create(c)

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, bookData)
	}
}

func (controller BookController) Update(c *gin.Context) {
	id := c.Params.ByName("id")

	var bookService service.BookService

	bookData, err := bookService.Update(id, c)

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, bookData)
	}
}

func (controller BookController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	var bookService service.BookService

	if err := bookService.Delete(id); err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
	} else {
		c.JSON(204, gin.H{"id": id})
	}
}
