package repository

import (
	"context"
	"errors"

	"github.com/matherique/share/internal/entity"
	"github.com/matherique/share/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *snipetsRepositoryMongo) Get(ctx context.Context, hash string, isSecure bool) (*entity.Snipet, error) {
	res := s.db.Collection(s.collection).FindOne(ctx, bson.M{"hash": hash, "is_secure": isSecure})

	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.ErrNotFound
		}

		return nil, err
	}

	var snip entity.Snipet
	if err := res.Decode(&snip); err != nil {
		return nil, err
	}

	return &snip, nil
}
