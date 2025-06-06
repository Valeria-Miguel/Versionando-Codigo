package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"tarea2/config"
	"tarea2/models"
)

func CreateTask(c *fiber.Ctx) error {
	userEmail := c.Locals("user").(string)
	userCol := config.GetCollection("users")
	taskCol := config.GetCollection("tasks")

	var user models.User
	err := userCol.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	task.ID = primitive.NewObjectID()
	task.UsuarioID = user.ID
	_, err = taskCol.InsertOne(context.TODO(), task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al crear tarea"})
	}

	return c.Status(201).JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	userEmail := c.Locals("user").(string)
	userCol := config.GetCollection("users")
	taskCol := config.GetCollection("tasks")

	var user models.User
	err := userCol.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Usuario no válido"})
	}

	cursor, err := taskCol.Find(context.TODO(), bson.M{"usuario_id": user.ID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener tareas"})
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al leer tareas"})
	}

	return c.JSON(tasks)
}

func UpdateTask(c *fiber.Ctx) error {
	taskCol := config.GetCollection("tasks")

	var payload struct {
		ID      string      `json:"id"`
		Updates models.Task `json:"updates"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	taskID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	update := bson.M{
		"$set": bson.M{
			"titulo":       payload.Updates.Titulo,
			"descripcion":  payload.Updates.Descripcion,
			"fecha_inicio": payload.Updates.FechaInicio,
			"deadline":     payload.Updates.Deadline,
		},
	}

	_, err = taskCol.UpdateOne(context.TODO(), bson.M{"_id": taskID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar tarea"})
	}

	return c.JSON(fiber.Map{"message": "Tarea actualizada"})
}

func DeleteTask(c *fiber.Ctx) error {
	taskCol := config.GetCollection("tasks")

	var payload struct {
		ID string `json:"id"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Petición inválida"})
	}

	taskID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	_, err = taskCol.DeleteOne(context.TODO(), bson.M{"_id": taskID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar tarea"})
	}

	return c.JSON(fiber.Map{"message": "Tarea eliminada"})
}


func GetTaskByID(c *fiber.Ctx) error {
	taskCol := config.GetCollection("tasks")

	var payload struct {
		ID string `json:"id"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Petición inválida"})
	}

	taskID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	var task models.Task
	err = taskCol.FindOne(context.TODO(), bson.M{"_id": taskID}).Decode(&task)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tarea no encontrada"})
	}

	return c.JSON(task)
}
