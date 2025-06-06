package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Task struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    UsuarioID   primitive.ObjectID `bson:"usuario_id,omitempty" json:"usuario_id,omitempty"`
    Titulo      string             `bson:"titulo" json:"titulo"`
    Descripcion string             `bson:"descripcion" json:"descripcion"`
    FechaInicio string             `bson:"fecha_inicio" json:"fecha_inicio"` 
    Deadline    string             `bson:"deadline" json:"deadline"`
}