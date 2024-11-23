package usecase

import (
	"context"
	"neuro-most/media-service/internal/entities"
	"time"
)

type (
	CreateMediaUseCase interface {
		Execute(ctx context.Context, input CreateMediaInput) error
	}
	CreateMediaInput struct {
		Title        string
		VideoURL     string
		ThumbnailURL string
		Description  string
	}

	createMediaInteractor struct {
		repo entities.MediaRepo
	}
)

func NewCreateMediaUseCase(repo entities.MediaRepo) CreateMediaUseCase {
	return &createMediaInteractor{repo: repo}
}

func (uc *createMediaInteractor) Execute(ctx context.Context, input CreateMediaInput) error {
	media := entities.NewMediaCreate(
		input.Title,
		input.VideoURL,
		input.ThumbnailURL,
		input.Description,
		"admin",
		time.Now(),
	)

	if err := uc.repo.Create(ctx, media); err != nil {
		return err
	}

	return nil
}
