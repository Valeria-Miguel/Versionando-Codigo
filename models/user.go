package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre           string             `json:"nombre"`
	Apellidos        string             `json:"apellidos"`
	Email            string             `json:"email"`
	Password         string             `json:"password"`
	FechaNacimiento  string             `json:"fecha_nacimiento"`
	PreguntaSecreta  string             `json:"pregunta_secreta"`
	RespuestaSecreta string             `json:"respuesta_secreta"`
}

