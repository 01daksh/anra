package internal

func InitializeTaskController() *TaskController {
	taskRepo := NewTaskRepository()
	taskServ := NewTaskService(taskRepo)
	taskController := NewTaskController(taskServ)

	return taskController
}
