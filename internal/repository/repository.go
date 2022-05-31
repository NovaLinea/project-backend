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
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error)
}

type User interface {
	GetDataProfile(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error)
	GetSettings(ctx context.Context, userID primitive.ObjectID) (domain.UserSettings, error)
	Save(ctx context.Context, userID primitive.ObjectID, inp domain.UserSaveSettings) error
	GetPasswordHash(ctx context.Context, userID primitive.ObjectID) (string, error)
	ChangePassword(ctx context.Context, userID primitive.ObjectID, password string) error
	DeleteAccount(ctx context.Context, userID primitive.ObjectID) error
	GetLikesFavorites(ctx context.Context, userID primitive.ObjectID) (domain.UserLikesFavorites, error)
	GetFollowsFollowings(ctx context.Context, userID primitive.ObjectID) (domain.UserFollowsFollowings, error)
	CheckSubscribe(ctx context.Context, fromID, toID primitive.ObjectID) (bool, error)
	Subscribe(ctx context.Context, userID, accoumtID primitive.ObjectID) error
	UnSubscribe(ctx context.Context, userID, accoumtID primitive.ObjectID) error
	GetFollows(ctx context.Context, userID primitive.ObjectID) ([]domain.UserProfile, error)
	GetFollowings(ctx context.Context, userID primitive.ObjectID) ([]domain.UserProfile, error)
}

type Project interface {
	GetPopular(ctx context.Context) ([]domain.ProjectData, error)
	GetDataProject(ctx context.Context, projectID primitive.ObjectID) (domain.ProjectData, error)

	Create(ctx context.Context, inp domain.ProjectData) error
	GetProjectsUser(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error)
	GetHome(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error)
	GetFavorites(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error)
	Like(ctx context.Context, projectID, userID primitive.ObjectID) error
	Dislike(ctx context.Context, projectID, userID primitive.ObjectID) error
	Favorite(ctx context.Context, projectID, userID primitive.ObjectID) error
	RemoveFavorite(ctx context.Context, projectID, userID primitive.ObjectID) error
	DeleteProject(ctx context.Context, projectID primitive.ObjectID) error
	Update(ctx context.Context, inp domain.ProjectEdit) error
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
