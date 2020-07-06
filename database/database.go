package database

import (
	"fmt"
	"mygo/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dialect := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, dbName)

	db, err = gorm.Open("mysql", dialect)

	if err != nil {
		panic(err)
	}

	AutoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func AutoMigration() {
	db.AutoMigrate(&model.User{}, &model.Book{})
}
