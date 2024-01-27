package entity

import (
	"context"
	"time"
)

var now = time.Now

type SnipetsRepository interface {
	Save(ctx context.Context, snipet *Snipet) error
	Get(ctx context.Context, hash string) (*Snipet, error)
}

type Snipet struct {
	Hash_link  string    `json:"hash_link" bson:"hash_link"`
	Content    string    `json:"content" bson:"content"`
	ExpirestAt time.Time `json:"expires_at" bson:"expires_at"`
}

func NewSnipet(hash, content string, duration int) *Snipet {
	return &Snipet{
		Hash_link:  hash,
		Content:    content,
		ExpirestAt: time.Now().AddDate(0, 0, duration),
	}
}
