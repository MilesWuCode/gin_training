package main

import (
	"mygo/database"
	"mygo/router"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.SetOutput(file)

	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	database.Init()

	router.Init()

	database.Close()
}
