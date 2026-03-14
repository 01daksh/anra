package interfaces

import (
	"anra/enums"
	"anra/models"
)

type TaskRepositoryInterface interface {
	AddTask(task models.Task) error
	GetTasks() ([]models.Task, error)
}

type TaskServiceInterface interface {
	GetTasks() ([]models.Task, error)
	AddTask(req enums.CreateRequest) (*models.Task, error)
}
