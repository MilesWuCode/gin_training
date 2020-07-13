package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name   string `gorm:"not null" form:"name" json:"name" binding:"required"`
	UserID uint   `gorm:"not null" form:"user_id" json:"user_id" binding:"required"`
}
