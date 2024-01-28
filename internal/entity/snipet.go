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
	HashLink   string    `json:"hash" bson:"hash"`
	Content    string    `json:"content" bson:"content"`
	ExpirestAt time.Time `json:"expires_at" bson:"expires_at"`
}

func NewSnipet(hash, content string, duration int) *Snipet {
	return &Snipet{
		HashLink:   hash,
		Content:    content,
		ExpirestAt: time.Now().AddDate(0, 0, duration),
	}
}
