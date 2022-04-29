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

func (r *UserRepo) GetData(ctx context.Context, userID primitive.ObjectID) (domain.UserData, error) {
	var data domain.UserData

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&data); err != nil {
		return domain.UserData{}, err
	}

	return data, nil
}

func (r *UserRepo) SaveData(ctx context.Context, userID primitive.ObjectID, inp domain.UserData) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"name": inp.Name, "email": inp.Email, "description": inp.Description, "phone": inp.Phone}})
	return err
}
