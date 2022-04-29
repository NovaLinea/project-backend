package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectData struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"user_id"`
	Name          string             `json:"name" binding:"required"`
	Descritpion   string             `json:"description" binding:"required"`
	Photo         string             `json:"photo"`
	Price         uint64             `json:"price"`
	PaymentSystem string             `json:"payment_system"`
	CountStaff    uint64             `json:"count_staff"`
	Staff         string             `json:"staff"`
	Progress      uint32             `json:"progress"`
	Type          string             `json:"type" binding:"required"`
	Time          string             `json:"time" binding:"required"`
}
