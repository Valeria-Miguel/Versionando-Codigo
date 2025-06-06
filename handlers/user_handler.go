package handlers

import (
	"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"tarea2/config"
	"tarea2/models"
	"tarea2/utils"
)

func Register(c *fiber.Ctx) error {
	userCol := config.GetCollection("users") 

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Petición inválida"})
	}

	user.Email = strings.ToLower(user.Email)

	// Verificar si ya existe
	count, _ := userCol.CountDocuments(context.TODO(), bson.M{"email": user.Email})
	if count > 0 {
		return c.Status(409).JSON(fiber.Map{"error": "El correo ya está registrado"})
	}

	// Hashear la contraseña
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashed)

	user.ID = primitive.NewObjectID()
	_, err := userCol.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Error al registrar usuario"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Usuario registrado"})
}

func Login(c *fiber.Ctx) error {
	userCol := config.GetCollection("users") 

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Petición inválida"})
	}

	var user models.User
	err := userCol.FindOne(context.TODO(), bson.M{"email": strings.ToLower(input.Email)}).Decode(&user)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	// Comparar contraseñas
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	// Generar JWT
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error generando token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// con token

func GetUsers(c *fiber.Ctx) error {
	userCol := config.GetCollection("users")

	cursor, err := userCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener usuarios"})
	}
	defer cursor.Close(context.TODO())

	var users []models.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al leer usuarios"})
	}

	// eliminar contraseñas del json
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	userCol := config.GetCollection("users")

	var payload struct {
		ID string `json:"id"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Petición inválida"})
	}

	userID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	var user models.User
	err = userCol.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}

	user.Password = ""
	return c.JSON(user)
}


func UpdateUser(c *fiber.Ctx) error {
	userCol := config.GetCollection("users")

	var payload map[string]interface{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	idStr, ok := payload["id"].(string)
	if !ok || idStr == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID faltante o inválido"})
	}

	userID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	delete(payload, "id")

	if passRaw, ok := payload["password"].(string); ok && passRaw != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(passRaw), 14)
		payload["password"] = string(hashed)
	}

	update := bson.M{
		"$set": payload,
	}

	_, err = userCol.UpdateOne(context.TODO(), bson.M{"_id": userID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al actualizar usuario"})
	}

	return c.JSON(fiber.Map{"message": "Usuario actualizado correctamente"})
}



func DeleteUser(c *fiber.Ctx) error {
	userCol := config.GetCollection("users")

	var payload struct {
		ID string `json:"id"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Petición inválida"})
	}

	userID, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}

	_, err = userCol.DeleteOne(context.TODO(), bson.M{"_id": userID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar usuario"})
	}

	return c.JSON(fiber.Map{"message": "Usuario eliminado"})
}

