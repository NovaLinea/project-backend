package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAuth struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}
