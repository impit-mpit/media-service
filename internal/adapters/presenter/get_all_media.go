package presenter

import (
	"neuro-most/media-service/internal/entities"
	"neuro-most/media-service/internal/usecase"
)

type getAllMediaPresenter struct{}

func NewGetAllMediaPresenter() getAllMediaPresenter {
	return getAllMediaPresenter{}
}

func (p getAllMediaPresenter) Output(media []entities.Media) []usecase.GetAllMediaOutput {
	var output []usecase.GetAllMediaOutput
	for _, m := range media {
		output = append(output, usecase.GetAllMediaOutput{
			ID:           m.ID(),
			Title:        m.Title(),
			VideoURL:     m.VideoURL(),
			ThumbnailURL: m.ThumbnailURL(),
			Description:  m.Description(),
		})
	}

	return output
}
