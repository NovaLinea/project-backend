package repository

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{db: db.Collection(usersCollection)}
}

func (r *UserRepo) GetDataProfile(ctx context.Context, userID primitive.ObjectID) (domain.UserProfile, error) {
	var data domain.UserProfile

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&data); err != nil {
		return domain.UserProfile{}, err
	}

	return data, nil
}

func (r *UserRepo) GetDataSettings(ctx context.Context, userID primitive.ObjectID) (domain.UserSettings, error) {
	var data domain.UserSettings

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&data); err != nil {
		return domain.UserSettings{}, err
	}

	return data, nil
}

func (r *UserRepo) SaveData(ctx context.Context, userID primitive.ObjectID, inp domain.UserSettings) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"name": inp.Name, "email": inp.Email, "description": inp.Description}})
	return err
}

func (r *UserRepo) GetPasswordHash(ctx context.Context, userID primitive.ObjectID) (string, error) {
	var data domain.UserAuth

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&data); err != nil {
		return "", err
	}

	return data.Password, nil
}

func (r *UserRepo) ChangePassword(ctx context.Context, userID primitive.ObjectID, password string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"password": password}})
	return err
}

func (r *UserRepo) DeleteAccount(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": userID})
	return err
}
