package service

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/ProjectUnion/project-backend.git/internal/repository"
	"github.com/ProjectUnion/project-backend.git/pkg/logging"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(repo repository.Project) *ProjectService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetPopular(ctx context.Context) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetPopular(ctx)
	return projects, err
}

func (s *ProjectService) GetDataProject(ctx context.Context, projectID primitive.ObjectID) (domain.ProjectData, error) {
	data, err := s.repo.GetDataProject(ctx, projectID)
	return data, err
}

func (s *ProjectService) Create(ctx context.Context, inp domain.ProjectData) error {
	err := s.repo.Create(ctx, inp)
	return err
}

func (s *ProjectService) GetProjectsUser(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetProjectsUser(ctx, userID)
	return projects, err
}

func (s *ProjectService) GetHome(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetHome(ctx, userID)
	return projects, err
}

func (s *ProjectService) GetFavorites(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetFavorites(ctx, userID)
	return projects, err
}

func (s *ProjectService) Like(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.Like(ctx, projectID, userID)
	return err
}

func (s *ProjectService) Dislike(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.Dislike(ctx, projectID, userID)
	return err
}

func (s *ProjectService) Favorite(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.Favorite(ctx, projectID, userID)
	return err
}

func (s *ProjectService) RemoveFavorite(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.RemoveFavorite(ctx, projectID, userID)
	return err
}

func (s *ProjectService) DeleteProject(ctx context.Context, projectID primitive.ObjectID) error {
	err := s.repo.DeleteProject(ctx, projectID)
	return err
}

func (s *ProjectService) Update(ctx context.Context, inp domain.ProjectEdit) error {
	err := s.repo.Update(ctx, inp)
	return err
}
