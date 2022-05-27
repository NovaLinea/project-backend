package service

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/ProjectUnion/project-backend.git/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userData struct {
	AccessToken string
	//RefreshToken string
	UserID string
}

type Authorization interface {
	Register(ctx context.Context, inp domain.UserAuth) (userData, error)
	Login(ctx context.Context, email, password string) (userData, error)
	ParseToken(token string) (string, error)
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error)
}

type User interface {
	GetDataProfile(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error)
	GetDataSettings(ctx context.Context, userID primitive.ObjectID) (domain.UserSettings, error)
	SaveData(ctx context.Context, userID primitive.ObjectID, inp domain.UserSettings) error
	ChangePassword(ctx context.Context, userID primitive.ObjectID, inp domain.ChangePassword) error
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

type Service struct {
	Authorization
	User
	Project
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
		User:          NewUserService(repos.User),
		Project:       NewProjectService(repos.Project),
	}
}
