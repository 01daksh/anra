package internal

import (
	"anra/enums"
	"anra/interfaces"
	"anra/models"
	"errors"
	"github.com/google/uuid"
	"strings"
)

type TaskService struct {
	Repo interfaces.TaskRepositoryInterface
}

func NewTaskService(repo interfaces.TaskRepositoryInterface) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (ts *TaskService) GetTasks() ([]models.Task, error) {
	return ts.Repo.GetTasks()
}

func (ts *TaskService) AddTask(req enums.CreateRequest) (*models.Task, error) {
	if len(req.Title) > 200 || strings.TrimSpace(req.Title) == "" {
		return nil, errors.New("title can not be empty and greater than 200 characters")
	}

	if req.Status == "" {
		req.Status = enums.Todo
	}

	if !enums.IsStatusValid(req.Status) {
		return nil, errors.New("invalid status")
	}

	task := models.Task{
		Title:  req.Title,
		Status: req.Status,
		Id:     uuid.New().String(),
	}

	err := ts.Repo.AddTask(task)
	if err != nil {
		return nil, errors.New("add task failed")
	}

	return &task, nil
}
