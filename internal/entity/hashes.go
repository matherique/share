package entity

import "context"

type HashesRepository interface {
	GetAvaliable(ctx context.Context) (string, error)
	IsAvaliable(ctx context.Context, hash string) (bool, error)
}

type Hashes struct {
	Hash        string `bson:"hash"`
	IsAvaliable bool   `bson:"is_avaliable"`
}
