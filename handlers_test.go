package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/takama/router"

	"github.com/BatON-RW/dummy-service/storage"
)

type StorageStub struct {
	task *storage.Task	
	err error
}

func (s *StorageStub) GetTask(name string)	(*storage.Task, error) {
	return s.task, s.err
}

func (s *StorageStub) GetAllTasks() (*[]storage.Task, error) {
	tasks := make([]storage.Task, 1)
	tasks[0] = *s.task
	return &tasks, s.err
}

func (s *StorageStub) AddTask(task *storage.Task) error {
	s.task = task
	return s.err
}

// TestHandler is the simplest test: check base (/) URL
func TestGetTaskHandler(t *testing.T) {
	tsk := &storage.Task{Name: "test1", ExpDate: "2017-10-08", Priority: 2, Comment: "test task"}
	storageStub := &StorageStub{task: tsk, err: nil}

	r := router.New()
	r.GET("/task/", getTaskHandler(storageStub))

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

	expectedTask := "{\"name\":\"test1\",\"exp_date\":\"2017-10-08\",\"priority\":2,\"comment\":\"test task\"}"
	testingTask := strings.Trim(string(task), " \n")
	if testingTask != expectedTask {
		t.Fatalf(
			"Wrong task '%s', expected '%s'",
			testingTask, expectedTask,
		)
	}
}

func TestGetAllTasksHandler(t *testing.T) {
	tsk := &storage.Task{Name: "test1", ExpDate: "2017-10-08", Priority: 2, Comment: "test task"}
	storageStub := &StorageStub{task: tsk, err: nil}

	r := router.New()
	r.GET("/tasks/", getAllTasksHandler(storageStub))

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

	expectedTasks := "[{\"name\":\"test1\",\"exp_date\":\"2017-10-08\",\"priority\":2,\"comment\":\"test task\"}]"
	testingTasks := strings.Trim(string(tasks), " \n")
	if testingTasks != expectedTasks {
		t.Fatalf(
			"Wrong task '%s', expected '%s'",
			testingTasks, expectedTasks,
		)
	}
}

func TestAddTaskHandler(t *testing.T) {
	storageStub := &StorageStub{task: nil, err: nil}
	
	r := router.New()
	r.POST("/task/", addTaskHandler(storageStub))

	ts := httptest.NewServer(r)
	defer ts.Close()

	inTask := "{\"name\":\"test1\",\"exp_date\":\"2017-10-08\",\"priority\":2,\"comment\":\"test task\"}"
	_, err := http.Post(ts.URL + "/task/", "application/json", strings.NewReader(inTask))
	if err != nil {
		t.Fatal(err)
	}

	expectedTask := &storage.Task{Name: "test1", ExpDate: "2017-10-08", Priority: 2, Comment: "test task"}
	testingTask := storageStub.task
	if *testingTask != *expectedTask {
		t.Fatalf(
			"Wrong task '%s', expected '%s'",
			testingTask, expectedTask,
		)
	}
}