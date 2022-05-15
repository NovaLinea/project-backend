package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectData struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"user_id" bson:"userid"`
	NameCreator   string             `json:"name_creator" bson:"namecreator"`
	PhotoCreator  string             `json:"photo_creator" bson:"photocreator"`
	Name          string             `json:"name" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Photos        []string           `json:"photos"`
	Price         uint64             `json:"price"`
	PaymentSystem string             `json:"payment_system" bson:"paymentsystem"`
	Staff         []string           `json:"staff"`
	Progress      uint32             `json:"progress"`
	Likes         uint32             `json:"likes"`
	Comments      []string           `json:"comments"`
	Type          string             `json:"type" binding:"required"`
	Time          time.Time          `json:"time" binding:"required"`
	EditingTime   time.Time          `json:"editing_time" bson:"editingtime"`
}

type ProjectEdit struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Photos        []string           `json:"photos"`
	Price         uint64             `json:"price"`
	PaymentSystem string             `json:"payment_system" bson:"paymentsystem"`
	Staff         []string           `json:"staff"`
	EditingTime   time.Time          `json:"editing_time" bson:"editingtime"`
}
