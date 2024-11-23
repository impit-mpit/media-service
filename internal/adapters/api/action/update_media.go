package action

import (
	"context"
	mediav1 "neuro-most/media-service/gen/go/media/v1"
	"neuro-most/media-service/internal/usecase"
)

type UpdateMediaAction struct {
	uc usecase.UpdateMediaUseCase
}

func NewUpdateMediaAction(uc usecase.UpdateMediaUseCase) *UpdateMediaAction {
	return &UpdateMediaAction{
		uc: uc,
	}
}

func (a *UpdateMediaAction) Execute(ctx context.Context, input *mediav1.UpdateMediaRequest) error {
	var usecaseInput usecase.UpdateMediaInput
	usecaseInput.ID = input.Id
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
