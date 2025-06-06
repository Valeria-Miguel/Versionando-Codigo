package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	uri := os.Getenv("MONGO_URI")
	log.Println("Conectando a MongoDB...")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error al conectar:", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Error al verificar conexión:", err)
	}

	DB = client.Database("fiber_api")
	log.Println("Conexión exitosa")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}