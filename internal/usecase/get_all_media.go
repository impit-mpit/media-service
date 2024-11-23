package usecase

import (
	"context"
	"neuro-most/media-service/internal/entities"
)

type (
	GetAllMediaUseCase interface {
		Execute(ctx context.Context, input GetAllMediaInput) ([]GetAllMediaOutput, int64, error)
	}

	GetAllMediaInput struct {
		Page     int
		PageSize int
	}

	GetAllMediaOutput struct {
		ID           int64
		Title        string
		VideoURL     string
		ThumbnailURL string
		Description  string
	}

	GetAllMediaPresenter interface {
		Output(media []entities.Media) []GetAllMediaOutput
	}

	getAllMediaInteractor struct {
		repo      entities.MediaRepo
		presenter GetAllMediaPresenter
	}
)

func NewGetAllMediaUseCase(repo entities.MediaRepo, presenter GetAllMediaPresenter) GetAllMediaUseCase {
	return &getAllMediaInteractor{repo: repo, presenter: presenter}
}

func (g *getAllMediaInteractor) Execute(ctx context.Context, input GetAllMediaInput) ([]GetAllMediaOutput, int64, error) {
	media, total, err := g.repo.Fetch(ctx, int64(input.Page), int64(input.PageSize))
	if err != nil {
		return nil, 0, err
	}

	return g.presenter.Output(media), total, nil
}
