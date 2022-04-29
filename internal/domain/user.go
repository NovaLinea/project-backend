package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAuth struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

type UserProfile struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description"`
}

type UserSettings struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required"`
	Email       string             `json:"email" binding:"required"`
	Description string             `json:"description"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
