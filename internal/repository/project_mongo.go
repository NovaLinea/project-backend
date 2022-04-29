package repository

import (
	"context"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
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
