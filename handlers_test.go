package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/takama/router"
)

// TestHandler is the simplest test: check base (/) URL
func TestGetTaskHandler(t *testing.T) {
	r := router.New()
	r.GET("/task/", getTaskHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/task/")
	if err != nil {
		t.Fatal(err)
	}

	task, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedTask := "{\"Name\":\"test1\",\"ExpDate\":\"2017-10-08\",\"Priority\":2,\"Comment\":\"test task\"}"
	testingTask := strings.Trim(string(task), " \n")
	if testingTask != expectedTask {
		t.Fatalf(
			"Wrong task '%s', expected '%s'",
			testingTask, expectedTask,
		)
	}
}

func TestGetAllTasksHandler(t *testing.T) {
	r := router.New()
	r.GET("/tasks/", getAllTasksHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/tasks/")
	if err != nil {
		t.Fatal(err)
	}

	tasks, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedTasks := "[{\"Name\":\"some task\",\"ExpDate\":\"2017-10-08\",\"Priority\":2,\"Comment\":\"test task\"}]"
	testingTasks := strings.Trim(string(tasks), " \n")
	if testingTasks != expectedTasks {
		t.Fatalf(
			"Wrong task '%s', expected '%s'",
			testingTasks, expectedTasks,
		)
	}
}