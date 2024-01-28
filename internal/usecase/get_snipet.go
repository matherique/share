package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/matherique/share/internal/entity"
	"github.com/matherique/share/pkg/utils"
)

var (
	ErrNotFound           = errors.New("snipet not found")
	ErrSomethingWentWrong = errors.New("ops! somthing went wrong")
)

type GetSnipetUseCase interface {
	Execute(ctx context.Context, h string) (*entity.Snipet, error)
}

type getSnipetUseCase struct {
	snipetRepository entity.SnipetsRepository
}

func NewGetSnipetUseCase(snipetRepository entity.SnipetsRepository) GetSnipetUseCase {
	return &getSnipetUseCase{
		snipetRepository: snipetRepository,
	}
}

func (g *getSnipetUseCase) Execute(ctx context.Context, h string) (*entity.Snipet, error) {
	s, err := g.snipetRepository.Get(ctx, h)

	if errors.Is(err, utils.ErrNotFound) {
		slog.Error(ErrNotFound.Error(), "err", err)
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return s, nil
}
