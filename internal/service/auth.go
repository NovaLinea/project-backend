package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/ProjectUnion/project-backend.git/internal/repository"
	"github.com/ProjectUnion/project-backend.git/pkg/logging"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	accessTokenTTL  = 30 * time.Minute
	refreshTokenTTL = 30 * 24 * time.Hour
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}

	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) Register(ctx context.Context, inp domain.UserAuth) (userData, error) {
	err := s.repo.CreateUser(ctx, inp.Name, inp.Email, generatePasswordHash(inp.Password))
	if err != nil {
		return userData{}, err
	}

	user, err := s.repo.GetUser(ctx, inp.Email, generatePasswordHash(inp.Password))
	if err != nil {
		return userData{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *AuthorizationService) Login(ctx context.Context, email, password string) (userData, error) {
	user, err := s.repo.GetUser(ctx, email, generatePasswordHash(password))
	if err != nil {
		return userData{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *AuthorizationService) Refresh(ctx context.Context, refreshToken string) (userData, error) {
	_, err := s.ParseToken(refreshToken)
	if err != nil {
		return userData{}, err
	}

	user, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return userData{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *AuthorizationService) Logout(ctx context.Context, refreshToken string) error {
	err := s.repo.RemoveRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthorizationService) CreateSession(ctx context.Context, userId primitive.ObjectID) (userData, error) {
	var (
		res userData
		err error
	)

	res.UserID = userId.Hex()
	res.AccessToken, err = NewJWT(userId.Hex(), accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = NewJWT(userId.Hex(), refreshTokenTTL)
	if err != nil {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(refreshTokenTTL),
	}

	err = s.repo.SetSession(ctx, userId, session)
	return res, err
}

func NewJWT(userId string, tokenTTL time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *AuthorizationService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
