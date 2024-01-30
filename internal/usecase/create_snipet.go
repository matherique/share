package usecase

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/matherique/share/internal/entity"
)

var (
	ErrUploadingFile = errors.New("error uploading file")
)

type CreateUseCase interface {
	Execute(ctx context.Context, r io.Reader, size int64) (string, error)
}

type createUseCase struct {
	maxSizeAllowed      int64
	snipetRepository    entity.SnipetsRepository
	generateHashUseCase GenerateHashUseCase
}

func NewCreateUseCase(hashrepository entity.HashesRepository, snipetrepository entity.SnipetsRepository) CreateUseCase {
	return &createUseCase{
		maxSizeAllowed:   1024 * 1024,
		snipetRepository: snipetrepository,
	}
}

func (a createUseCase) Execute(ctx context.Context, r io.Reader, size int64) (string, error) {
	buff := new(bytes.Buffer)

	n, err := io.CopyN(buff, r, size)
	if err != nil {
		slog.Error(ErrUploadingFile.Error(), "err", err)
		return "", ErrUploadingFile
	}

	slog.Info("receive bytes", "size", n)

	link, err := a.generateHashUseCase.Execute(ctx)
	if err != nil {
		return "", err
	}

	snipet := entity.NewSnipet(link, buff.String(), 1)

	go func() {
		if err := a.snipetRepository.Save(context.Background(), snipet); err != nil {
			slog.Error("fail on save snipet", "err", err)
		}
	}()

	return link, nil
}
