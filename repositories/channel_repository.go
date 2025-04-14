package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type ChannelRepository interface {
	GetAllChannel(ctx context.Context) ([]entities.Channel, error)
	GetChannelById(ctx context.Context, channelId int) (entities.Channel, error)
	CreateChannel(ctx context.Context, channel entities.Channel) (entities.Channel, error)
	UpdateChannel(ctx context.Context, channel entities.Channel) (entities.Channel, error)
	DeleteChannel(ctx context.Context, channelId int) error
}

type channelRepository struct {
	db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	return &channelRepository{
		db: db,
	}
}

func (r *channelRepository) GetAllChannel(ctx context.Context) ([]entities.Channel, error) {
	var channels []entities.Channel
	err := r.db.Find(&channels).Error
	if err != nil {
		return []entities.Channel{}, err
	}
	return channels, err
}

func (r *channelRepository) GetChannelById(ctx context.Context, channelId int) (entities.Channel, error) {
	var channel entities.Channel
	err := r.db.Where("id = ?", channelId).First(&channel).Error
	if err != nil {
		return entities.Channel{}, err
	}
	return channel, err
}

func (r *channelRepository) CreateChannel(ctx context.Context, channel entities.Channel) (entities.Channel, error) {
	err := r.db.Create(&channel).Error
	if err != nil {
		return entities.Channel{}, err
	}
	return channel, err
}

func (r *channelRepository) UpdateChannel(ctx context.Context, channel entities.Channel) (entities.Channel, error) {
	err := r.db.Save(&channel).Error
	if err != nil {
		return entities.Channel{}, err
	}
	return channel, err
}

func (r *channelRepository) DeleteChannel(ctx context.Context, channelId int) error {
	err := r.db.Delete(&entities.Channel{}, "id = ?", channelId).Error
	if err != nil {
		return err
	}
	return nil
}
