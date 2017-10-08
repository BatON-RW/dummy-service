package main

type Task struct {
	Name string
	ExpDate string
	Priority int
	Comment string
}


func getTask(name string) (*Task, error) {
	return &Task{Name: name, ExpDate: "2017-10-08", Priority: 2, Comment: "test task"}, nil
}

func getAllTasks() (*[]Task, error) {
	tasks := make([]Task, 1)
	tasks[0] = Task{Name: "some task", ExpDate: "2017-10-08", Priority: 2, Comment: "test task"}
	return &tasks, nil
}