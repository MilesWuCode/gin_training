package main

import (
	"mygo/database"
	"mygo/router"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// json
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	// text
	logrus.SetFormatter(&logrus.TextFormatter{})

	file, _ := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logrus.SetOutput(file)

	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	database.Init()

	router.Init()

	database.Close()
}
