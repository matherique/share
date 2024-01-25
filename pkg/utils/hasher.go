package utils

import (
	"encoding/base64"
	"hash"
)

type Hasher func(b hash.Hash) string

func GenerateRandomHash(h hash.Hash) string {
	base := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return base[:8]
}
