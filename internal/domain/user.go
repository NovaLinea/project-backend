package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAuth struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}

type UserCreate struct {
	ID            primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name          string               `json:"name" binding:"required"`
	Email         string               `json:"email" binding:"required"`
	Password      string               `json:"password" binding:"required"`
	Description   string               `json:"description"`
	VerifyEmail   bool                 `json:"verify_email"`
	Photo         string               `json:"photo"`
	Follows       []primitive.ObjectID `json:"follows"`
	Followings    []primitive.ObjectID `json:"followings"`
	Favorites     []primitive.ObjectID `json:"favorites"`
	Likes         []primitive.ObjectID `json:"likes"`
	Notifications TypeNotifications    `json:"notifications"`
}

type UserProfile struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description"`
	VerifyEmail bool               `json:"verify_email"`
	Photo       string             `json:"photo"`
	Follows     []string           `json:"follows"`
	Followings  []string           `json:"followings"`
	Favorites   []string           `json:"favorites"`
}

type UserSettings struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" binding:"required"`
	Email         string             `json:"email" binding:"required"`
	VerifyEmail   bool               `json:"verify_email"`
	Description   string             `json:"description"`
	Notifications TypeNotifications  `json:"notifications"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type TypeNotifications struct {
	NewMessage        bool `json:"new_message" bson:"new_message"`
	NewSub            bool `json:"new_sub" bson:"new_sub"`
	NewComment        bool `json:"new_comment" bson:"new_comment"`
	Update            bool `json:"update" bson:"update"`
	EmailNotification bool `json:"email_notification" bson:"email_notification"`
}
