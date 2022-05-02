package service

import (
	"context"
	"errors"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/ProjectUnion/project-backend.git/internal/repository"
	"github.com/ProjectUnion/project-backend.git/pkg/logging"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	return &UserService{repo: repo}
}

func (s *UserService) GetDataProfile(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error) {
	data, err := s.repo.GetDataProfile(ctx, userID)
	return data, err
}

func (s *UserService) GetDataSettings(ctx context.Context, userID primitive.ObjectID) (domain.UserSettings, error) {
	data, err := s.repo.GetDataSettings(ctx, userID)
	return data, err
}

func (s *UserService) SaveData(ctx context.Context, userID primitive.ObjectID, inp domain.UserSettings) error {
	err := s.repo.SaveData(ctx, userID, inp)
	return err
}

func (s *UserService) ChangePassword(ctx context.Context, userID primitive.ObjectID, inp domain.ChangePassword) error {
	password, err := s.repo.GetPasswordHash(ctx, userID)
	if err != nil {
		return err
	}

	if password != generatePasswordHash(inp.OldPassword) {
		return errors.New("Incorrect old password")
	} else {
		if err = s.repo.ChangePassword(ctx, userID, generatePasswordHash(inp.NewPassword)); err != nil {
			return err
		}
	}

	return nil
}

func (s *UserService) DeleteAccount(ctx context.Context, userID primitive.ObjectID) error {
	err := s.repo.DeleteAccount(ctx, userID)
	return err
}

func (s *UserService) GetLikesFavorites(ctx context.Context, userID primitive.ObjectID) (domain.UserLikesFavorites, error) {
	lists, err := s.repo.GetLikesFavorites(ctx, userID)
	return lists, err
}
