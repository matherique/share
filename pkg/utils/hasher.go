package utils

import (
	"crypto/rand"
	"encoding/base64"
)

type Hasher func() string

func GenerateRandomHash() string {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	base := base64.StdEncoding.EncodeToString(b)
	return base[:8]
}
