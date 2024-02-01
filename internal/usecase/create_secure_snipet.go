package usecase

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/matherique/share/internal/entity"
	"github.com/matherique/share/pkg/secure"
)

var (
	ErrFailToEncrypt = errors.New("fail to encrypt")
)

type CreateSecureSnipetUseCase interface {
	Execute(ctx context.Context, r io.Reader, size int64) (string, string, error)
}

type createSecureUseCase struct {
	maxSizeAllowed      int64
	snipetRepository    entity.SnipetsRepository
	generateHashUseCase GenerateHashUseCase
	secure              secure.Secure
}

func NewCreateSecureUseCase(snipetrepository entity.SnipetsRepository, generageHash GenerateHashUseCase, encrypt secure.Secure) CreateSecureSnipetUseCase {
	return &createSecureUseCase{
		maxSizeAllowed:      1024 * 1024,
		snipetRepository:    snipetrepository,
		generateHashUseCase: generageHash,
		secure:              encrypt,
	}
}

func (c *createSecureUseCase) Execute(ctx context.Context, r io.Reader, size int64) (string, string, error) {
	buff := new(bytes.Buffer)

	n, err := io.CopyN(buff, r, size)
	if err != nil {
		slog.Error(ErrUploadingFile.Error(), "err", err)
		return "", "", ErrUploadingFile
	}

	slog.Info("receive bytes", "size", n)

	key, text, err := c.secure.Encrypt(buff.Bytes())

	if err != nil {
		slog.Error(ErrFailToEncrypt.Error(), "err", err)
		return "", "", ErrFailToEncrypt
	}

	link, err := c.generateHashUseCase.Execute(ctx)
	if err != nil {
		return "", "", err
	}

	snipet := entity.NewSnipet(link, text, 1)
	snipet.SetSecure()

	go func() {
		if err := c.snipetRepository.Save(context.Background(), snipet); err != nil {
			slog.Error("fail on save snipet", "err", err)
		}
	}()

	return key, link, nil
}
