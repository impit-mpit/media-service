package usecase

import (
	"context"
	"neuro-most/media-service/internal/entities"
)

type (
	DeleteMediaUseCase interface {
		Execute(ctx context.Context, input DeleteMediaInput) error
	}

	DeleteMediaInput struct {
		ID int64
	}

	deleteMediaInteractor struct {
		repo entities.MediaRepo
	}
)

func NewDeleteMediaUseCase(repo entities.MediaRepo) DeleteMediaUseCase {
	return &deleteMediaInteractor{repo: repo}
}

func (uc *deleteMediaInteractor) Execute(ctx context.Context, input DeleteMediaInput) error {
	media, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	if err := uc.repo.Delete(ctx, media.ID()); err != nil {
		return err
	}

	return nil
}
