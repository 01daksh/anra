package main

import (
	"anra/internal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	runHttpRouter()
}

func runHttpRouter() {
	app := fiber.New(fiber.Config{
		AppName: "Anra Tasks",
	})

	//register endpoints
	// enabled logger for incoming request
	app.Use(logger.New())

	//get task list
	app.Get("/tasks", internal.InitializeTaskController().GetTasks)
	//add task (post)
	app.Post("/tasks", internal.InitializeTaskController().AddTask)

	err := app.Listen(":3001")
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
