package repository

import (
	"context"
	"errors"

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

func (r *AuthorizationRepo) CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error) {
	var user domain.UserProfile

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserProfile{}, domain.ErrUserNotFound
		}
		return domain.UserProfile{}, err
	}

	return user, nil
}
