package tests

import (
	"anra/enums"
	"anra/internal"
	"testing"
)

func TestCreateTaskSuccess(t *testing.T) {
	taskController := internal.InitializeTaskController()

	req := enums.CreateRequest{
		Title:  "Wake up Early",
		Status: "",
	}

	task, err := taskController.Service.AddTask(req)
	if err != nil {
		t.Fatalf("Error creating task: %v", err)
	}

	if task.Status != "todo" {
		t.Fatalf("Error interpreting empty task status: %s", task.Status)
	}

}

func TestCreateTaskInvalidStatus(t *testing.T) {
	taskController := internal.InitializeTaskController()

	req := enums.CreateRequest{
		Title:  "Learn Go",
		Status: "invalid_status",
	}

	_, err := taskController.Service.AddTask(req)

	if err == nil {
		t.Fatalf("Expected error for invalid status but got nil")
	}
}

func TestCreateTaskTitleTooLong(t *testing.T) {
	taskController := internal.InitializeTaskController()

	longTitle := ""
	for i := 0; i < 20; i++ {
		longTitle = longTitle + "lorem ipsum"
	}

	req := enums.CreateRequest{
		Title:  longTitle,
		Status: enums.Todo,
	}

	_, err := taskController.Service.AddTask(req)

	if err == nil {
		t.Fatalf("Expected error for title longer than 200 characters")
	}
}

func TestCreateTaskTitleRequired(t *testing.T) {
	taskController := internal.InitializeTaskController()

	req := enums.CreateRequest{
		Title:  "",
		Status: enums.Todo,
	}

	_, err := taskController.Service.AddTask(req)

	if err == nil {
		t.Fatalf("Expected error when title is empty")
	}
}

func TestGetTasks(t *testing.T) {
	taskController := internal.InitializeTaskController()

	req := enums.CreateRequest{
		Title:  "Task 1",
		Status: enums.Todo,
	}

	//add task
	_, err := taskController.Service.AddTask(req)
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	// get task
	tasks, err := taskController.Service.GetTasks()
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	// verify if task was added
	if len(tasks) == 0 || tasks[0].Title != "Task 1" {
		t.Fatalf("unexpected task title: %s", tasks[0].Title)
	}
}
