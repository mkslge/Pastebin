package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("bins"),
	}
}

func (r *Repository) Create(ctx context.Context, bin Bin) error {
	_, err := r.collection.InsertOne(ctx, bin)
	return err
}
