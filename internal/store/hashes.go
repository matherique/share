package store

import (
	"context"
	"errors"
	"time"

	"github.com/matherique/share/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	now = time.Now
)

type HashesStore interface {
	GetAvaliable(ctx context.Context) (string, error)
	IsAvaliable(ctx context.Context, hash string) (bool, error)
}

type hashesStore struct {
	db         *mongo.Database
	collection string
}

func NewHashesStore(db *mongo.Database) HashesStore {
	return &hashesStore{
		db:         db,
		collection: "hashes",
	}
}

func (h hashesStore) IsAvaliable(ctx context.Context, hash string) (bool, error) {
	filter := bson.M{
		"hash": hash,
	}

	res := h.db.Collection(h.collection).FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}

		return false, err
	}

	var hashes entity.Hashes

	if err := res.Decode(&hashes); err != nil {
		return false, nil
	}

	return hashes.IsAvaliable, nil
}

func (h hashesStore) GetAvaliable(ctx context.Context) (string, error) {
	filter := bson.M{
		"is_avaliable": true,
	}

	update := bson.M{
		"$set": bson.M{
			"is_avaliable": false,
		},
	}

	opts := options.FindOneAndUpdate().SetSort(bson.M{"created_at": 1})

	res := h.db.Collection(h.collection).FindOneAndUpdate(ctx, filter, update, opts)
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", nil
		}

		return "", err
	}

	var hashes entity.Hashes

	if err := res.Decode(&hashes); err != nil {
		return "", err
	}

	return hashes.Hash, nil
}
