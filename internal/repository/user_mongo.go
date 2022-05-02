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
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$set": bson.M{"name": inp.Name, "email": inp.Email, "description": inp.Description, "notifications": inp.Notifications}})
	if err != nil {
		return err
	}

	_, err = r.db.Database().Collection(projectsCollection).UpdateOne(ctx, bson.M{"userid": userID}, bson.M{"$set": bson.M{"namecreator": inp.Name}})
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
	if err != nil {
		return err
	}

	_, err = r.db.Database().Collection(projectsCollection).DeleteMany(ctx, bson.M{"userid": userID})
	return err
}

func (r *UserRepo) GetLikesFavorites(ctx context.Context, userID primitive.ObjectID) (domain.UserLikesFavorites, error) {
	var lists domain.UserLikesFavorites

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&lists); err != nil {
		return domain.UserLikesFavorites{}, err
	}

	return lists, nil
}

func (r *UserRepo) GetDataParams(ctx context.Context, userID primitive.ObjectID) (domain.UserDataParams, error) {
	var data domain.UserDataParams

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&data); err != nil {
		return domain.UserDataParams{}, err
	}

	return data, nil
}

func (r *UserRepo) SubscribeUser(ctx context.Context, userID, accoumtID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"followings": accoumtID}})
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{"_id": accoumtID}, bson.M{"$push": bson.M{"follows": userID}})
	return err
}

func (r *UserRepo) UnSubscribeUser(ctx context.Context, userID, accoumtID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"followings": accoumtID}})
	if err != nil {
		return err
	}

	_, err = r.db.UpdateOne(ctx, bson.M{"_id": accoumtID}, bson.M{"$pull": bson.M{"follows": userID}})
	return err
}
