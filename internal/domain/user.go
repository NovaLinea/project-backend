package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSignUp struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" binding:"required"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

type UserSignIn struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}
