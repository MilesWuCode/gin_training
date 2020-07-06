package service

import (
	"mygo/database"
	"mygo/model"

	"github.com/gin-gonic/gin"
)

type BookService struct{}

type Book model.Book

func (service BookService) GetAll() ([]Book, error) {
	db := database.GetDB()

	var bookData []Book

	if err := db.Find(&bookData).Error; err != nil {
		return nil, err
	}

	return bookData, nil
}

func (service BookService) Get(id string) (Book, error) {
	db := database.GetDB()

	var bookData Book

	if err := db.Where("id = ?", id).First(&bookData).Error; err != nil {
		return bookData, err
	}

	return bookData, nil
}

func (service BookService) Create(c *gin.Context) (Book, error) {
	db := database.GetDB()

	var bookData Book

	if err := c.Bind(&bookData); err != nil {
		return bookData, err
	}

	if err := db.Create(&bookData).Error; err != nil {
		return bookData, err
	}

	return bookData, nil
}

func (service BookService) Update(id string, c *gin.Context) (Book, error) {
	db := database.GetDB()

	var bookData Book

	if err := db.Where("id = ?", id).First(&bookData).Error; err != nil {
		return bookData, err
	}

	if err := c.Bind(&bookData); err != nil {
		return bookData, err
	}

	db.Save(&bookData)

	return bookData, nil
}

func (service BookService) Delete(id string) error {
	db := database.GetDB()

	var bookData Book

	if err := db.Where("id = ?", id).Delete(&bookData).Error; err != nil {
		return err
	}

	return nil
}
