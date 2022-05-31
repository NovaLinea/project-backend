package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectData struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatorID     primitive.ObjectID `json:"creator_id" bson:"userid"`
	CreatorName   string             `json:"creator_name" bson:"namecreator"`
	CreatorPhoto  string             `json:"creator_photo" bson:"photocreator"`
	Name          string             `json:"name" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Photos        []string           `json:"photos"`
	Price         uint64             `json:"price"`
	PaymentSystem string             `json:"payment_system" bson:"paymentsystem"`
	Staff         []string           `json:"staff"`
	Progress      uint32             `json:"progress"`
	Views         uint32             `json:"views"`
	Likes         uint32             `json:"likes"`
	Type          string             `json:"type" binding:"required"`
	CreatedAt     time.Time          `json:"created_at" bson:"createat"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updatedat"`
}

type ProjectEdit struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Photos        []string           `json:"photos"`
	Price         uint64             `json:"price"`
	PaymentSystem string             `json:"payment_system" bson:"paymentsystem"`
	Staff         []string           `json:"staff"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updatedat"`
}
