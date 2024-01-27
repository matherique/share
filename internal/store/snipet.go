package store

import (
	"context"

	"github.com/matherique/share/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type SnipetsStore interface {
	Save(ctx context.Context, snipet *entity.Snipet) error
}

type snipetsStore struct {
	db         *mongo.Database
	collection string
}

func NewSnipetStore(db *mongo.Database) SnipetsStore {
	return &snipetsStore{
		db:         db,
		collection: "snipets",
	}
}

func (s snipetsStore) Save(ctx context.Context, snipet *entity.Snipet) error {
	_, err := s.db.Collection(s.collection).InsertOne(ctx, snipet)
	if err != nil {
		return err
	}

	return nil
}
