package usecase

import (
	"context"
	"neuro-most/media-service/internal/entities"
)

type (
	GetByIDMediaUseCase interface {
		Execute(ctx context.Context, input GetByIDMediaInput) (GetByIDMediaOutput, error)
	}

	GetByIDMediaInput struct {
		ID int64
	}

	GetByIDMediaOutput struct {
		ID           int64
		Title        string
		VideoURL     string
		ThumbnailURL string
		Description  string
	}

	GetByIDMediaPresenter interface {
		Output(media entities.Media) GetByIDMediaOutput
	}

	getByIDMediaInteractor struct {
		repo      entities.MediaRepo
		presenter GetByIDMediaPresenter
	}
)

func NewGetByIDMediaUseCase(repo entities.MediaRepo, presenter GetByIDMediaPresenter) GetByIDMediaUseCase {
	return &getByIDMediaInteractor{repo: repo, presenter: presenter}
}

func (uc *getByIDMediaInteractor) Execute(ctx context.Context, input GetByIDMediaInput) (GetByIDMediaOutput, error) {
	media, err := uc.repo.GetByID(ctx, input.ID)
	if err != nil {
		return GetByIDMediaOutput{}, err
	}

	return uc.presenter.Output(media), nil
}
