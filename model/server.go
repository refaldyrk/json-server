package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Server struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	ServerID  string             `json:"server_id" bson:"server_id"`
	Data      []byte             `json:"data" bson:"data"`
	Raw       map[string]any     `json:"raw" bson:"raw"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
}
