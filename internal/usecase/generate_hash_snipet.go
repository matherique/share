package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/matherique/share/internal/entity"
	"github.com/matherique/share/pkg/utils"
)

var (
	ErrCheckingIFLinkExists = errors.New("error checking if link already exists")
	ErrGetAvaliableLink     = errors.New("error getting avaliable link")
)

type GenerateHashUseCase interface {
	Execute(ctx context.Context) (string, error)
}

type generateHashUseCase struct {
	hasher         utils.Hasher
	hashRepository entity.HashesRepository
}

func NewGenerateHashUseCase(hashRepository entity.HashesRepository) *generateHashUseCase {
	return &generateHashUseCase{
		hasher:         utils.GenerateRandomHash,
		hashRepository: hashRepository,
	}
}

func (g *generateHashUseCase) Execute(ctx context.Context) (string, error) {
	link := g.hasher()

	has, err := g.hashRepository.IsAvaliable(ctx, link)

	if err != nil {
		slog.Error(ErrCheckingIFLinkExists.Error(), "err", err)
		return "", ErrCheckingIFLinkExists
	}

	if !has {
		link, err = g.hashRepository.GetAvaliable(ctx)

		if err != nil {
			slog.Error(ErrGetAvaliableLink.Error(), "err", err)
			return "", ErrGetAvaliableLink
		}
	}

	return link, nil
}
