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

func (s *ProjectService) CreateProject(ctx context.Context, inp domain.ProjectData) error {
	err := s.repo.CreateProject(ctx, inp)
	return err
}

func (s *ProjectService) GetProjectsPopular(ctx context.Context) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetProjectsPopular(ctx)
	return projects, err
}

func (s *ProjectService) GetProjectsUser(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetProjectsUser(ctx, userID)
	return projects, err
}

func (s *ProjectService) GetProjectsHome(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetProjectsHome(ctx, userID)
	return projects, err
}

func (s *ProjectService) GetFavoritesProjects(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	projects, err := s.repo.GetFavoritesProjects(ctx, userID)
	return projects, err
}

func (s *ProjectService) LikeProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.LikeProject(ctx, projectID, userID)
	return err
}

func (s *ProjectService) DislikeProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.DislikeProject(ctx, projectID, userID)
	return err
}

func (s *ProjectService) FavoriteProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.FavoriteProject(ctx, projectID, userID)
	return err
}

func (s *ProjectService) RemoveFavoriteProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	err := s.repo.RemoveFavoriteProject(ctx, projectID, userID)
	return err
}

func (s *ProjectService) GetDataProject(ctx context.Context, projectID primitive.ObjectID) (domain.ProjectData, error) {
	data, err := s.repo.GetDataProject(ctx, projectID)
	return data, err
}
