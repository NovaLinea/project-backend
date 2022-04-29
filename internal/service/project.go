package service

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/ProjectUnion/project-backend.git/internal/repository"
	"github.com/ProjectUnion/project-backend.git/pkg/logging"
	"github.com/joho/godotenv"
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
