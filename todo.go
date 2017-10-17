package main

import (
	"github.com/takama/router"
	"github.com/Sirupsen/logrus"
	"os"

	"github.com/BatON-RW/dummy-service/storage"
)

var log = logrus.New()

func main() {
	port := Getenv("SERVICE_PORT", "8000")

	conf := &storage.StorageConfig{
		Host: Getenv("MONGO_PORT_27017_TCP_ADDR", "some-mongo"),
		Port: Getenv("MONGO_PORT_27017_TCP_PORT", "27017"),
		DB: "todolist",
		Collection: "test"}
	storage, err := storage.New(conf)
	if err != nil {
		log.Errorf("Storage initialisation error: %s", err)
		panic(err)
	}
	defer storage.C.Database.Session.Close()

	r := router.New()
	r.GET("/task/:name", getTaskHandler(storage))
	r.POST("/task/", addTaskHandler(storage))
	r.GET("/tasks/", getAllTasksHandler(storage))
	r.GET("/version/", versionHandler)
	r.Logger = logger
	r.Listen("0.0.0.0:" + port)
}

// Getenv extracts ENV variable if exist, else returns deafault value.
func Getenv(key string, deflt string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		env = deflt
	}
	return env
}