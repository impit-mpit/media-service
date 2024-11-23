package action

import (
	"context"
	mediav1 "neuro-most/media-service/gen/go/media/v1"
	"neuro-most/media-service/internal/usecase"
)

type GetByIDMediaAction struct {
	uc usecase.GetByIDMediaUseCase
}

func NewGetByIDMediaAction(uc usecase.GetByIDMediaUseCase) *GetByIDMediaAction {
	return &GetByIDMediaAction{
		uc: uc,
	}
}

func (a *GetByIDMediaAction) Execute(ctx context.Context, input *mediav1.GetMediaByIdRequest) (*mediav1.Media, error) {
	var usecaseInput usecase.GetByIDMediaInput
	usecaseInput.ID = input.Id

	output, err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return nil, err
	}

	return &mediav1.Media{
		Id:           output.ID,
		Title:        output.Title,
		VideoUrl:     output.VideoURL,
		ThumbnailUrl: output.ThumbnailURL,
		Description:  output.Description,
	}, nil
}
