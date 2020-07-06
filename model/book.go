package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name string `gorm:"unique;not null" form:"name" json:"name" binding:"required"`
}
