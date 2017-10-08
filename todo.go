package main

import (
	"github.com/takama/router"
	"github.com/Sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		port = "8000"
	}

	r := router.New()
	r.GET("/task/", getTaskHandler)
	r.GET("/tasks/", getAllTasksHandler)
	r.GET("/version/", versionHandler)
	r.Logger = logger
	r.Listen("0.0.0.0:" + port)
}