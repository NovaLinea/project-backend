package repository

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectRepo struct {
	db *mongo.Collection
}

func NewProjectRepo(db *mongo.Database) *ProjectRepo {
	return &ProjectRepo{db: db.Collection(projectsCollection)}
}

func (r *ProjectRepo) CreateProject(ctx context.Context, inp domain.ProjectData) error {
	inp.Comments = []string{}

	_, err := r.db.InsertOne(ctx, inp)
	return err
}

func (r *ProjectRepo) GetProjects(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	var projects []domain.ProjectData

	filter := bson.M{"userid": userID}
	sort := bson.M{"time": -1}

	opts := options.FindOptions{
		Sort: &sort,
	}

	cur, err := r.db.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepo) GetFavoritesProjects(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	var projects []domain.ProjectData
	var user domain.UserLikesFavorites

	if err := r.db.Database().Collection(usersCollection).FindOne(ctx, bson.M{"_id": userID}).Decode(&user); err != nil {
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$in": user.Favorites}}
	sort := bson.M{"time": -1}

	opts := options.FindOptions{
		Sort: &sort,
	}

	cur, err := r.db.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepo) LikeProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$inc": bson.M{"likes": 1}})
	if err != nil {
		return err
	}

	_, err = r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"likes": projectID}})
	return err
}

func (r *ProjectRepo) DislikeProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$inc": bson.M{"likes": -1}})
	if err != nil {
		return err
	}

	_, err = r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"likes": projectID}})
	return err
}

func (r *ProjectRepo) FavoriteProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"favorites": projectID}})
	return err
}

func (r *ProjectRepo) RemoveFavoriteProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"favorites": projectID}})
	return err
}
