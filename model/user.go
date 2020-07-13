package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" form:"name" json:"name" binding:"required"`
	Email    string `gorm:"unique;not null" form:"email" json:"email" binding:"required"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required"`
	Books    []Book `gorm:"foreignkey:UserID"`
}
