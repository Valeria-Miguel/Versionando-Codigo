package routes

import (
	"github.com/gofiber/fiber/v2"
	"tarea2/handlers"
	"tarea2/middleware"
)

func TaskRoutes(app *fiber.App) {
	task := app.Group("/api/tasks", middleware.JWTProtected())

	task.Post("/create", handlers.CreateTask)
	task.Get("/tasks", handlers.GetTasks)
	task.Post("/get", handlers.GetTaskByID)
	task.Put("/update", handlers.UpdateTask)
	task.Delete("/delete", handlers.DeleteTask)
}
