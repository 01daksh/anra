package internal

import (
	"anra/models"
	"sync"
)

type TaskRepository struct {
	InMemoryTasks *models.TaskCollection
	mu            sync.Mutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		InMemoryTasks: models.NewTaskCollection(),
	}
}

func (tr *TaskRepository) AddTask(task models.Task) error {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	tr.InMemoryTasks.AddTask(task)
	return nil
}

func (tr *TaskRepository) GetTasks() ([]models.Task, error) {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	return tr.InMemoryTasks.GetTasks(), nil
}
