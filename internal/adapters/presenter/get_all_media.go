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
	const maxLength = 40

	if len(description) == 0 {
		return ""
	}

	runes := []rune(description)

	if len(runes) <= maxLength {
		return description
	}

	lastSpace := -1
	for i := 0; i < maxLength && i < len(runes); i++ {
		if runes[i] == ' ' {
			lastSpace = i
		}
	}

	cutoff := maxLength
	if lastSpace != -1 {
		cutoff = lastSpace
	}

	return string(runes[:cutoff]) + "..."
}
