package repository

import (
	"context"

	"github.com/matherique/share/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type snipetsRepositoryMongo struct {
	db         *mongo.Database
	collection string
}

func NewSnipetRepositoryMongo(db *mongo.Database) entity.SnipetsRepository {
	return &snipetsRepositoryMongo{
		db:         db,
		collection: "snipets",
	}
}

func (s snipetsRepositoryMongo) Save(ctx context.Context, snipet *entity.Snipet) error {
	_, err := s.db.Collection(s.collection).InsertOne(ctx, snipet)
	if err != nil {
		return err
	}

	return nil
}
