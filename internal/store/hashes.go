package store

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
)

type HashesStore interface {
	Get(ctx context.Context, hash string) (string, error)
}

type hashesStore struct {
	db *sqlx.DB
}

func NewHashesStore(db *sqlx.DB) HashesStore {
	return &hashesStore{
		db: db,
	}
}

func (h hashesStore) Get(ctx context.Context, hash string) (string, error) {
	return "", errors.New("not implemented")
}
