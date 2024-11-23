package entities

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrMediaNotFound = status.New(codes.NotFound, "media not found").Err()
	ErrorMediaCreate = status.New(codes.Internal, "error create media").Err()
	ErrorMediaUpdate = status.New(codes.Internal, "error update media").Err()
	ErrorMediaDelete = status.New(codes.Internal, "error delete media").Err()
	ErrorMediaFetch  = status.New(codes.Internal, "error fetch news").Err()
)

type (
	MediaRepo interface {
		Create(ctx context.Context, media Media) error
		Update(ctx context.Context, media Media) error
		Delete(ctx context.Context, id int64) error
		Fetch(ctx context.Context, page, pageSize int64) ([]Media, int64, error)
		GetByID(ctx context.Context, id int64) (Media, error)
	}

	Media struct {
		id           int64
		title        string
		videoURL     string
		thumbnailURL string
		description  string
		createdDate  time.Time
		createdBy    string
		updatedDate  *time.Time
		updatedBy    *string
	}
)

func NewMedia(
	id int64,
	title string,
	videoURL string,
	thumbnailURL string,
	description string,
	createdDate time.Time,
	createdBy string,
	updatedDate *time.Time,
	updatedBy *string,
) Media {
	return Media{
		id:           id,
		title:        title,
		videoURL:     videoURL,
		thumbnailURL: thumbnailURL,
		description:  description,
		createdDate:  createdDate,
		createdBy:    createdBy,
		updatedDate:  updatedDate,
		updatedBy:    updatedBy,
	}
}

func NewMediaCreate(
	title string,
	videoURL string,
	thumbnailURL string,
	description string,
	createdBy string,
	createdDate time.Time,
) Media {
	return Media{
		title:        title,
		videoURL:     videoURL,
		thumbnailURL: thumbnailURL,
		description:  description,
		createdBy:    createdBy,
		createdDate:  createdDate,
	}
}

func (m Media) ID() int64 {
	return m.id
}

func (m Media) Title() string {
	return m.title
}

func (m Media) VideoURL() string {
	return m.videoURL
}

func (m Media) ThumbnailURL() string {
	return m.thumbnailURL
}

func (m Media) Description() string {
	return m.description
}

func (m Media) CreatedBy() string {
	return m.createdBy
}

func (m Media) UpdatedBy() *string {
	return m.updatedBy
}

func (m Media) CreatedDate() time.Time {
	return m.createdDate
}

func (m Media) UpdatedDate() *time.Time {
	return m.updatedDate
}

func (m *Media) SetID(id int64) {
	m.id = id
}

func (m *Media) SetTitle(title string) {
	m.title = title
}

func (m *Media) SetVideoURL(videoURL string) {
	m.videoURL = videoURL
}

func (m *Media) SetThumbnailURL(thumbnailURL string) {
	m.thumbnailURL = thumbnailURL
}

func (m *Media) SetDescription(description string) {
	m.description = description
}

func (m *Media) SetCreatedBy(createdBy string) {
	m.createdBy = createdBy
}

func (m *Media) SetUpdatedBy(updatedBy *string) {
	m.updatedBy = updatedBy
}

func (m *Media) SetCreatedDate(createdDate time.Time) {
	m.createdDate = createdDate
}

func (m *Media) SetUpdatedDate(updatedDate *time.Time) {
	m.updatedDate = updatedDate
}
