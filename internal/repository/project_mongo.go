package repository

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepo struct {
	db *mongo.Collection
}

func NewProjectRepo(db *mongo.Database) *ProjectRepo {
	return &ProjectRepo{db: db.Collection(projectsCollection)}
}

func (r *ProjectRepo) CreateProject(ctx context.Context, inp domain.ProjectData) error {
	_, err := r.db.InsertOne(ctx, inp)
	return err
}

func (r *ProjectRepo) GetProjects(ctx context.Context, userID primitive.ObjectID) ([]domain.ProjectData, error) {
	var projects []domain.ProjectData

	cur, err := r.db.Find(ctx, bson.M{"userid": userID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}