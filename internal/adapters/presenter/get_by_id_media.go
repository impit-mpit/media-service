package presenter

import (
	"neuro-most/media-service/internal/entities"
	"neuro-most/media-service/internal/usecase"
)

type getByIdMediaPresenter struct{}

func NewGetByIDMediaPresenter() getByIdMediaPresenter {
	return getByIdMediaPresenter{}
}

func (p getByIdMediaPresenter) Output(media entities.Media) usecase.GetByIDMediaOutput {
	return usecase.GetByIDMediaOutput{
		ID:           media.ID(),
		Title:        media.Title(),
		VideoURL:     media.VideoURL(),
		ThumbnailURL: media.ThumbnailURL(),
		Description:  media.Description(),
	}
}
