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
			//short descirption 40 symbols
			ShortDescription: ShortDescription(m.Description()),
		})
	}

	return output
}

func ShortDescription(description string) string {
	if len(description) > 40 {
		return description[:40]
	}
	return description
}
