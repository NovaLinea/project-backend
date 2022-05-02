package repository

import (
	"context"
	"errors"
	"time"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorizationRepo struct {
	db *mongo.Collection
}

func NewAuthorizationRepo(db *mongo.Database) *AuthorizationRepo {
	return &AuthorizationRepo{db: db.Collection(usersCollection)}
}

func (r *AuthorizationRepo) CreateUser(ctx context.Context, name, email, password string) error {
	var inp domain.UserCreate
	var ntfs domain.TypeNotifications

	ntfs.NewMessage = true
	ntfs.NewComment = true
	ntfs.NewSub = true
	ntfs.Update = true
	ntfs.EmailNotification = true

	inp.Name = name
	inp.Email = email
	inp.Password = password
	inp.Favorites = []primitive.ObjectID{}
	inp.Follows = []primitive.ObjectID{}
	inp.Followings = []primitive.ObjectID{}
	inp.Likes = []primitive.ObjectID{}
	inp.Notifications = ntfs

	_, err := r.db.InsertOne(ctx, inp)
	return err
}

func (r *AuthorizationRepo) GetUser(ctx context.Context, email, password string) (domain.UserAuth, error) {
	var user domain.UserAuth

	if err := r.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserAuth{}, domain.ErrUserNotFound
		}
		return domain.UserAuth{}, err
	}

	return user, nil
}

func (r *AuthorizationRepo) SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{
		"$set": bson.M{
			"session": session,
		}})

	return err
}

func (r *AuthorizationRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.UserAuth, error) {
	var user domain.UserAuth
	if err := r.db.FindOne(ctx, bson.M{
		"session.refreshToken": refreshToken,
		"session.expiresAt":    bson.M{"$gt": time.Now()},
	}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserAuth{}, domain.ErrUserNotFound
		}
		return domain.UserAuth{}, err

	}

	return user, nil
}

func (r *AuthorizationRepo) RemoveRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"session.refreshToken": refreshToken}, bson.M{
		"$set": bson.M{
			"session.refreshToken": "",
		}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFound
		}
		return err
	}

	return nil
}
