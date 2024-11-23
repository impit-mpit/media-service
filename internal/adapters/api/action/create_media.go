package action

import (
	"context"
	mediav1 "neuro-most/media-service/gen/go/media/v1"
	"neuro-most/media-service/internal/usecase"
)

type CreateMediaAction struct {
	uc usecase.CreateMediaUseCase
}

func NewCreateMediaAction(uc usecase.CreateMediaUseCase) *CreateMediaAction {
	return &CreateMediaAction{
		uc: uc,
	}
}

func (a *CreateMediaAction) Execute(ctx context.Context, input *mediav1.CreateMediaRequest) error {
	var usecaseInput usecase.CreateMediaInput
	usecaseInput.Title = input.Title
	usecaseInput.VideoURL = input.VideoUrl
	usecaseInput.ThumbnailURL = input.ThumbnailUrl
	usecaseInput.Description = input.Description

	err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return err
	}
	return nil
}
