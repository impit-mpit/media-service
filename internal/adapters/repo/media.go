package repo

import (
	"context"
	"fmt"
	"neuro-most/media-service/internal/entities"
	"time"
)

type mediaGORM struct {
	ID           int64     `gorm:"primaryKey"`
	Title        string    `gorm:"not null"`
	VideoURL     string    `gorm:"not null"`
	ThumbnailURL string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	CreatedDate  time.Time `gorm:"not null"`
	CreatedBy    string    `gorm:"not null"`
	UpdatedDate  *string
	UpdatedBy    *time.Time
}

type MediaRepo struct {
	db GSQL
}

func NewMediaRepo(db GSQL) MediaRepo {
	db.AutoMigrate(&mediaGORM{})
	return MediaRepo{db: db}
}

func (m MediaRepo) Create(ctx context.Context, media entities.Media) error {
	mediaGORM := mediaGORM{
		Title:        media.Title(),
		VideoURL:     media.VideoURL(),
		ThumbnailURL: media.ThumbnailURL(),
		Description:  media.Description(),
		CreatedDate:  media.CreatedDate(),
		CreatedBy:    media.CreatedBy(),
	}
	if err := m.db.Create(ctx, &mediaGORM); err != nil {
		return entities.ErrorMediaCreate
	}
	return nil
}

func (m MediaRepo) Update(ctx context.Context, media entities.Media) error {
	updates := map[string]interface{}{
		"title":         media.Title(),
		"video_url":     media.VideoURL(),
		"thumbnail_url": media.ThumbnailURL(),
		"description":   media.Description(),
	}
	if err := m.db.UpdateOne(ctx, updates, &mediaGORM{ID: media.ID()}, &mediaGORM{}); err != nil {
		fmt.Println(err)
		return entities.ErrorMediaUpdate
	}
	return nil
}

func (m MediaRepo) Fetch(ctx context.Context, page, pageSize int64) ([]entities.Media, int64, error) {
	var medias []mediaGORM
	query := m.db.BeginFind(ctx, &medias)
	var total int64
	query.Count(&total)
	fmt.Println(total)
	query = query.Page(int(page), int(pageSize)).OrderBy("id desc")
	err := query.Find(&medias)
	if err != nil {
		return nil, 0, entities.ErrorMediaFetch
	}
	var result []entities.Media
	for _, media := range medias {
		result = append(result, m.convertMedia(media))
	}
	return result, total, nil
}

func (m MediaRepo) GetByID(ctx context.Context, id int64) (entities.Media, error) {
	var media mediaGORM
	if err := m.db.BeginFind(ctx, &media).Where(&mediaGORM{ID: id}).First(&media); err != nil {
		return entities.Media{}, entities.ErrMediaNotFound
	}
	return m.convertMedia(media), nil
}

func (m MediaRepo) Delete(ctx context.Context, id int64) error {
	if err := m.db.Delete(ctx, &mediaGORM{}, &mediaGORM{ID: id}); err != nil {
		return entities.ErrorMediaDelete
	}
	return nil
}

func (m MediaRepo) convertMedia(media mediaGORM) entities.Media {
	return entities.NewMedia(
		media.ID,
		media.Title,
		media.VideoURL,
		media.ThumbnailURL,
		media.Description,
		media.CreatedDate,
		media.CreatedBy,
		media.UpdatedBy,
		media.UpdatedDate,
	)
}
