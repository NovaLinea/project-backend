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
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error)
}

type User interface {
	GetDataProfile(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error)
	GetSettings(ctx context.Context, userID primitive.ObjectID) (domain.UserSettings, error)
	Save(ctx context.Context, userID primitive.ObjectID, inp domain.UserSaveSettings) error
	ChangePassword(ctx context.Context, userID primitive.ObjectID, inp domain.ChangePassword) error
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
