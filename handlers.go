package main

import (
    "fmt"
	"encoding/json"
	"net/http"

	"github.com/takama/router"

	"github.com/BatON-RW/dummy-service/version"
	"github.com/BatON-RW/dummy-service/storage"
)

const internalErrorText = "Internal server error."


func getTaskHandler(s storage.StorageIntf) func(c *router.Control) {
	return func(c *router.Control) {
		task, err := s.GetTask(c.Get(":name"))
		if err != nil {
			log.Errorf("Task getting error: %s", err)
			http.Error(c.Writer, internalErrorText, http.StatusInternalServerError)
			return
		}

		if mTask, err := json.Marshal(task); err == nil {
			fmt.Fprintf(c.Writer, string(mTask))
		} else {
			log.Errorf("JSON marshalling error: %s", err)
			http.Error(c.Writer, internalErrorText, http.StatusInternalServerError)
			return
		}
	}
}

func getAllTasksHandler(s storage.StorageIntf) func(c *router.Control) {
	return func(c *router.Control) {
		tasks, err := s.GetAllTasks()
		if err != nil {
			log.Errorf("Tasks getting error: %s", err)
			http.Error(c.Writer, internalErrorText, http.StatusInternalServerError)
			return
		}

		if mTask, err := json.Marshal(tasks); err == nil {
			fmt.Fprintf(c.Writer, string(mTask))
		} else {
			log.Errorf("JSON marshalling error: %s", err)
			http.Error(c.Writer, internalErrorText, http.StatusInternalServerError)
			return
		}
	}
}

func addTaskHandler(s storage.StorageIntf) func(c *router.Control) {
	return func(c *router.Control) {
		decoder := json.NewDecoder(c.Request.Body)
		defer c.Request.Body.Close()
		task := storage.Task{}
		err := decoder.Decode(&task)
		if err != nil {
			log.Errorf("JSON decoding error: %s", err)
			http.Error(c.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.AddTask(&task)
		if err != nil {
			log.Errorf("Task saving error: %s", err)
			http.Error(c.Writer, internalErrorText, http.StatusInternalServerError)
			return
		}
	}
}

func versionHandler(c *router.Control) {
	fmt.Fprintf(c.Writer, "Repo: %s, Commit: %s, Version: %s", version.REPO, version.COMMIT, version.RELEASE)	
}

// logger provides a log of requests
func logger(c *router.Control) {
	remoteAddr := c.Request.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = c.Request.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, c.Request.Method, c.Request.URL.Path)
}