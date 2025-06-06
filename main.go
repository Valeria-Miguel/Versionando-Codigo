package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"tarea2/config"
	"tarea2/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	app := fiber.New()

	config.ConnectDB()

	routes.UserRoutes(app)
	routes.TaskRoutes(app)
	app.Listen(":3000")
}
