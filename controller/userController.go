package controller

import (
	"mygo/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct{}

func (controller UserController) Index(c *gin.Context) {
	var userService service.UserService

	userData, err := userService.GetAll()

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(404, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, userData)
	}
}

func (controller UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")

	var userService service.UserService

	userData, err := userService.Get(id)

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(404, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, userData)
	}
}

func (controller UserController) Create(c *gin.Context) {
	var userService service.UserService

	userData, err := userService.Create(c)

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, userData)
	}
}

func (controller UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")

	var userService service.UserService

	userData, err := userService.Update(id, c)

	if err != nil {
		logrus.Warn(err.Error())

		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, userData)
	}
}

func (controller UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	var userService service.UserService

	if err := userService.Delete(id); err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
	} else {
		c.JSON(204, gin.H{"id": id})
	}
}
