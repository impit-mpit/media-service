package usecase

import (
	"context"
	"neuro-most/media-service/internal/entities"
)

type (
	UpdateMediaUseCase interface {
		Execute(ctx context.Context, input UpdateMediaInput) error
	}

	UpdateMediaInput struct {
		ID           int64
		Title        *string
		VideoURL     *string
		ThumbnailURL *string
		Description  *string
	}

	updateMediaInteractor struct {
		repo entities.MediaRepo
	}
)

func NewUpdateMediaUseCase(repo entities.MediaRepo) UpdateMediaUseCase {
	return &updateMediaInteractor{repo: repo}
}

func (u *updateMediaInteractor) Execute(ctx context.Context, input UpdateMediaInput) error {
	media, err := u.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if input.Title != nil {
		media.SetTitle(*input.Title)
	}

	if input.VideoURL != nil {
		media.SetVideoURL(*input.VideoURL)
	}

	if input.ThumbnailURL != nil {
		media.SetThumbnailURL(*input.ThumbnailURL)
	}

	if input.Description != nil {
		media.SetDescription(*input.Description)
	}

	if err := u.repo.Update(ctx, media); err != nil {
		return err
	}

	return nil
}
