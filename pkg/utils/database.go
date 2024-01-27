package utils

import "go.mongodb.org/mongo-driver/mongo"

var (
	ErrNotFound = mongo.ErrNoDocuments
)
