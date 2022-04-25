package service

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/ProjectUnion/project-backend.git/internal/repository"
)

type userData struct {
	AccessToken  string
	Position     string
	RefreshToken string
	UserID       string
}

type Authorization interface {
	Register(ctx context.Context, inp domain.UserAuth) (userData, error)
	Login(ctx context.Context, email, password string) (userData, error)
	Refresh(ctx context.Context, refreshToken string) (userData, error)
	Logout(ctx context.Context, refreshToken string) error
	ParseToken(token string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.Authorization),
	}
}
