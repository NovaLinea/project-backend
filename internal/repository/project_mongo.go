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

func (r *ProjectRepo) GetPopular(ctx context.Context) ([]domain.ProjectData, error) {
	var projects []domain.ProjectData

	filter := bson.M{}
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

func (r *ProjectRepo) GetDataProject(ctx context.Context, projectID primitive.ObjectID) (domain.ProjectData, error) {
	var data domain.ProjectData

	if err := r.db.FindOne(ctx, bson.M{"_id": projectID}).Decode(&data); err != nil {
		return domain.ProjectData{}, err
	}

	return data, nil
}

func (r *ProjectRepo) Create(ctx context.Context, inp domain.ProjectData) error {
	_, err := r.db.InsertOne(ctx, inp)
	return err
}

func (r *ProjectRepo) GetProjectsUser(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
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

func (r *ProjectRepo) GetHome(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	var projects []domain.ProjectData
	var user domain.UserFollowsFollowings

	if err := r.db.Database().Collection(usersCollection).FindOne(ctx, bson.M{"_id": userID}).Decode(&user); err != nil {
		return nil, err
	}

	filter := bson.M{"userid": bson.M{"$in": user.Followings}}
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

func (r *ProjectRepo) GetFavorites(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
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

func (r *ProjectRepo) Like(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$inc": bson.M{"likes": 1}})
	if err != nil {
		return err
	}

	_, err = r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"likes": projectID}})
	return err
}

func (r *ProjectRepo) Dislike(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$inc": bson.M{"likes": -1}})
	if err != nil {
		return err
	}

	_, err = r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"likes": projectID}})
	return err
}

func (r *ProjectRepo) Favorite(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$push": bson.M{"favorites": projectID}})
	return err
}

func (r *ProjectRepo) RemoveFavorite(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": userID}, bson.M{"$pull": bson.M{"favorites": projectID}})
	return err
}

func (r *ProjectRepo) DeleteProject(ctx context.Context, projectID primitive.ObjectID) error {
	var users []domain.UserID

	cur, err := r.db.Database().Collection(usersCollection).Find(ctx, bson.M{"likes": bson.M{"$all": []primitive.ObjectID{projectID}}})
	if err != nil {
		return err
	}

	if err := cur.All(ctx, &users); err != nil {
		return err
	}

	for _, user := range users {
		_, err = r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$pull": bson.M{"likes": projectID}})
		if err != nil {
			return err
		}
	}

	users = []domain.UserID{}
	cur, err = r.db.Database().Collection(usersCollection).Find(ctx, bson.M{"favorites": bson.M{"$all": []primitive.ObjectID{projectID}}})
	if err != nil {
		return err
	}

	if err := cur.All(ctx, &users); err != nil {
		return err
	}

	for _, user := range users {
		_, err = r.db.Database().Collection(usersCollection).UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$pull": bson.M{"favorites": projectID}})
		if err != nil {
			return err
		}
	}

	_, err = r.db.DeleteOne(ctx, bson.M{"_id": projectID})
	return err
}

func (r *ProjectRepo) Update(ctx context.Context, inp domain.ProjectEdit) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": inp.ID}, bson.M{
		"$set": bson.M{
			"name":          inp.Name,
			"description":   inp.Description,
			"price":         inp.Price,
			"paymentsystem": inp.PaymentSystem,
			"staff":         inp.Staff,
			"updatedat":     inp.UpdatedAt,
		},
	})
	return err
}
