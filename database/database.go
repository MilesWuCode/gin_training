package database

import (
	"fmt"
	"mygo/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	dbName := viper.GetString("db.database")
	userName := viper.GetString("db.username")
	password := viper.GetString("db.password")

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
