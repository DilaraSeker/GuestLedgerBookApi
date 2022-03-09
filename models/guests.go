package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guests struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Message string             `json:"message,omitempty"`
	Email   string             `json:"email,omitempty"`
}
