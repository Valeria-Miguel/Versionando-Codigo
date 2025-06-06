package routes

import (
	"tarea2/handlers"

	"github.com/gofiber/fiber/v2"
	"tarea2/middleware"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/api/users")
	user.Post("/register", handlers.Register)
	user.Post("/login", handlers.Login)
	
	// CRUD  con JWT

	user.Use(middleware.JWTProtected())
	user.Get("/users", handlers.GetUsers)
	user.Post("/get", handlers.GetUserByID)
	user.Put("/update", handlers.UpdateUser)
	user.Delete("/delete", handlers.DeleteUser)
}
