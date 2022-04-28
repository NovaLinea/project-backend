package repository

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(ctx context.Context, name, email, password string) error
	GetUser(ctx context.Context, email, password string) (domain.UserAuth, error)
	SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.UserAuth, error)
	RemoveRefreshToken(ctx context.Context, refreshToken string) error
}

type User interface {
	GetData(ctx context.Context, userID primitive.ObjectID) (domain.UserData, error)
}

type Repository struct {
	Authorization
	User
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepo(db.Database(viper.GetString("mongo.databaseName"))),
		User:          NewUserRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
