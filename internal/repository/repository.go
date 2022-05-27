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
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error)
}

type User interface {
	GetDataProfile(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error)
	GetDataSettings(ctx context.Context, userID primitive.ObjectID) (domain.UserSettings, error)
	SaveData(ctx context.Context, userID primitive.ObjectID, inp domain.UserSettings) error
	GetPasswordHash(ctx context.Context, userID primitive.ObjectID) (string, error)
	ChangePassword(ctx context.Context, userID primitive.ObjectID, password string) error
	DeleteAccount(ctx context.Context, userID primitive.ObjectID) error
	GetLikesFavorites(ctx context.Context, userID primitive.ObjectID) (domain.UserLikesFavorites, error)
	GetFollowsFollowings(ctx context.Context, userID primitive.ObjectID) (domain.UserFollowsFollowings, error)
	SubscribeUser(ctx context.Context, userID, accoumtID primitive.ObjectID) error
	UnSubscribeUser(ctx context.Context, userID, accoumtID primitive.ObjectID) error
	GetFollows(ctx context.Context, userID primitive.ObjectID) ([]domain.UserProfile, error)
	GetFollowings(ctx context.Context, userID primitive.ObjectID) ([]domain.UserProfile, error)
}

type Project interface {
	CreateProject(ctx context.Context, inp domain.ProjectData) error
	GetProjectsPopular(ctx context.Context) ([]domain.ProjectData, error)
	GetProjectsUser(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error)
	GetProjectsHome(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error)
	GetFavoritesProjects(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error)
	LikeProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	DislikeProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	FavoriteProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	RemoveFavoriteProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	GetDataProject(ctx context.Context, projectID primitive.ObjectID) (domain.ProjectData, error)
	DeleteProject(ctx context.Context, projectID primitive.ObjectID) error
	EditProject(ctx context.Context, inp domain.ProjectEdit) error
}

type Repository struct {
	Authorization
	User
	Project
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepo(db.Database(viper.GetString("mongo.databaseName"))),
		User:          NewUserRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Project:       NewProjectRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
