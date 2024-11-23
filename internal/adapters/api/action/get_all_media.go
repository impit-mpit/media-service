package action

import (
	"context"
	mediav1 "neuro-most/media-service/gen/go/media/v1"
	"neuro-most/media-service/internal/usecase"
)

type GetAllMediaAction struct {
	uc usecase.GetAllMediaUseCase
}

func NewGetAllMediaAction(uc usecase.GetAllMediaUseCase) *GetAllMediaAction {
	return &GetAllMediaAction{
		uc: uc,
	}
}

func (a *GetAllMediaAction) Execute(ctx context.Context, input *mediav1.GetMediaFeedRequest) (*mediav1.GetMediaFeedResponse, error) {
	var usecaseInput usecase.GetAllMediaInput
	usecaseInput.Page = int(input.Page)
	usecaseInput.PageSize = int(input.PageSize)
	media, total, err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return nil, err
	}
	mediaFeed := &mediav1.GetMediaFeedResponse{
		Total: int32(total),
		Media: make([]*mediav1.Media, 0),
	}
	for _, m := range media {
		mediaFeed.Media = append(mediaFeed.Media, &mediav1.Media{
			Id:           m.ID,
			Title:        m.Title,
			VideoUrl:     m.VideoURL,
			ThumbnailUrl: m.ThumbnailURL,
			Description:  m.Description,
		})
	}
	return mediaFeed, nil

}
