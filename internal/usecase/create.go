package usecase

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"hash"
	"io"
	"log/slog"

	"github.com/matherique/share/internal/entity"
	"github.com/matherique/share/internal/store"
	"github.com/matherique/share/pkg/utils"
)

var (
	ErrUploadingFile        = errors.New("error uploading file")
	ErrCheckingIFLinkExists = errors.New("error checking if link already exists")
	ErrGetAvaliableLink     = errors.New("error getting avaliable link")
)

type CreateUseCase interface {
	Execute(ctx context.Context, r io.Reader, size int64) (string, error)
}

type createUseCase struct {
	maxSizeAllowed int64
	hasher         utils.Hasher
	hashStore      store.HashesStore
	snipetStore    store.SnipetsStore
}

func NewCreateUseCase(hashStore store.HashesStore, snipetStore store.SnipetsStore) CreateUseCase {
	return &createUseCase{
		maxSizeAllowed: 1024 * 1024,
		hasher:         utils.GenerateRandomHash,
		hashStore:      hashStore,
		snipetStore:    snipetStore,
	}
}

func (a createUseCase) Execute(ctx context.Context, r io.Reader, size int64) (string, error) {
	buff := new(bytes.Buffer)

	hashSha256 := sha256.New()

	mw := io.MultiWriter(buff, hashSha256)

	n, err := io.CopyN(mw, r, size)
	if err != nil {
		slog.Error(ErrUploadingFile.Error(), "err", err)
		return "", ErrUploadingFile
	}

	slog.Info("receive bytes", "size", n)

	link, err := a.getLink(ctx, hashSha256)
	if err != nil {
		return "", err
	}

	snipet := entity.NewSnipet(link, buff.String(), 1)

	go func() {
		if err := a.snipetStore.Save(ctx, snipet); err != nil {
			slog.Error("fail on save snipet", "err", err)
		}
	}()

	return link, nil
}

func (a createUseCase) getLink(ctx context.Context, h hash.Hash) (string, error) {
	link := a.hasher(h)

	has, err := a.hashStore.IsAvaliable(ctx, link)

	if err != nil {
		slog.Error(ErrCheckingIFLinkExists.Error(), "err", err)
		return "", ErrCheckingIFLinkExists
	}

	if !has {
		link, err = a.hashStore.GetAvaliable(ctx)

		if err != nil {
			slog.Error(ErrGetAvaliableLink.Error(), "err", err)
			return "", ErrGetAvaliableLink
		}
	}

	return link, nil
}
