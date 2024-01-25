package usecase

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"io"
	"log/slog"

	"github.com/matherique/share/pkg/utils"
)

type ImportUseCase interface {
	Execute(ctx context.Context, r io.Reader, size int64) (string, error)
}

type importApp struct {
	maxSizeAllowed int64
	hasher         utils.Hasher
}

func NewImportUseCase() ImportUseCase {
	return &importApp{
		maxSizeAllowed: 1024 * 1024,
		hasher:         utils.GenerateRandomHash,
	}
}

func (a importApp) Execute(ctx context.Context, r io.Reader, size int64) (string, error) {
	buff := new(bytes.Buffer)

	hashSha256 := sha256.New()

	mw := io.MultiWriter(buff, hashSha256)

	n, err := io.CopyN(mw, r, size)
	if err != nil {
		slog.Error("error uploading file", "err", err)
		return "", errors.New("error uploading file: " + err.Error())
	}

	slog.Info("receive bytes", "size", n)
	link := a.hasher(hashSha256)

	return link, nil
}
