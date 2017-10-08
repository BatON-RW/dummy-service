package main

import (
    "fmt"
	"encoding/json"
	"net/http"

	"github.com/takama/router"

	"github.com/BatON-RW/dummy-service/version"
)


func getTaskHandler(c *router.Control) {
	task, _ := getTask("test1")
	if mTask, err := json.Marshal(task); err == nil {
		fmt.Fprintf(c.Writer, string(mTask))
	} else {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func getAllTasksHandler(c *router.Control) {
	tasks, _ := getAllTasks()
	if mTask, err := json.Marshal(tasks); err == nil {
		fmt.Fprintf(c.Writer, string(mTask))
	} else {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
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