package internal

import (
	"anra/enums"
	"anra/interfaces"
	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	Service interfaces.TaskServiceInterface
}

func NewTaskController(serv interfaces.TaskServiceInterface) *TaskController {
	return &TaskController{
		Service: serv,
	}
}

func (tc *TaskController) GetTasks(fCtx *fiber.Ctx) error {
	tasks, err := tc.Service.GetTasks()
	if err != nil {
		return err
	}

	return fCtx.JSON(tasks)
}

func (tc *TaskController) AddTask(fCtx *fiber.Ctx) error {
	var req enums.CreateRequest
	err := fCtx.BodyParser(&req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	task, err := tc.Service.AddTask(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fCtx.Status(fiber.StatusCreated).JSON(task)
}
