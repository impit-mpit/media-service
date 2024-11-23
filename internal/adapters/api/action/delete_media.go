package action

import (
	"context"
	mediav1 "neuro-most/media-service/gen/go/media/v1"
	"neuro-most/media-service/internal/usecase"
)

type DeleteMediaAction struct {
	uc usecase.DeleteMediaUseCase
}

func NewDeleteMediaAction(uc usecase.DeleteMediaUseCase) *DeleteMediaAction {
	return &DeleteMediaAction{
		uc: uc,
	}
}

func (a *DeleteMediaAction) Execute(ctx context.Context, input *mediav1.DeleteMediaRequest) error {
	var usecaseInput usecase.DeleteMediaInput
	usecaseInput.ID = input.Id

	err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return err
	}
	return nil
}
