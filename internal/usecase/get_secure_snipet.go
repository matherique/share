package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/matherique/share/internal/entity"
	"github.com/matherique/share/pkg/secure"
	"github.com/matherique/share/pkg/utils"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

type GetSecureSnipetUseCase interface {
	Execute(ctx context.Context, h string, key []byte) (*entity.Snipet, error)
}

type getSecureSnipetUseCase struct {
	snipetRepository entity.SnipetsRepository
	secure           secure.Secure
	isSecure         bool
}

func NewGetSecureSnipetUseCase(snipetRepo entity.SnipetsRepository, sec secure.Secure) *getSecureSnipetUseCase {
	return &getSecureSnipetUseCase{
		snipetRepository: snipetRepo,
		secure:           sec,
		isSecure:         true,
	}
}

func (g *getSecureSnipetUseCase) Execute(ctx context.Context, h string, key []byte) (*entity.Snipet, error) {
	s, err := g.snipetRepository.Get(ctx, h, g.isSecure)

	if errors.Is(err, utils.ErrNotFound) {
		slog.Error(ErrNotFound.Error(), "err", err)
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	msg, err := g.secure.Decrypt(key, s.Content)
	if err != nil {
		slog.Error(ErrInvalidKey.Error(), "err", err)
		return nil, ErrInvalidKey
	}

	s.Content = msg
	return s, nil
}
