package main

import (
	"mygo/database"
	"mygo/router"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	loadConfig()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.SetOutput(file)

	logrus.SetLevel(logrus.InfoLevel)
}

func loadConfig() {
	path, _ := os.Getwd()

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path)
	viper.ReadInConfig()
}

func main() {
	database.Init()

	router.Init()

	database.Close()
}
