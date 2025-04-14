package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type PlatformRepository interface {
	GetAllPlatform(ctx context.Context) ([]entities.Platform, error)
	GetPlatformById(ctx context.Context, platformId int) (entities.Platform, error)
	CreatePlatform(ctx context.Context, platform entities.Platform) (entities.Platform, error)
	UpdatePlatform(ctx context.Context, platform entities.Platform) (entities.Platform, error)
	DeletePlatform(ctx context.Context, platformId int) error
}

type platformRepository struct {
	db *gorm.DB
}

func NewPlatformRepository(db *gorm.DB) PlatformRepository {
	return &platformRepository{
		db: db,
	}
}

func (r *platformRepository) GetAllPlatform(ctx context.Context) ([]entities.Platform, error) {
	var platforms []entities.Platform
	err := r.db.Find(&platforms).Error
	if err != nil {
		return []entities.Platform{}, err
	}
	return platforms, err
}

func (r *platformRepository) GetPlatformById(ctx context.Context, platformId int) (entities.Platform, error) {
	var platform entities.Platform
	err := r.db.Where("id = ?", platformId).First(&platform).Error
	if err != nil {
		return entities.Platform{}, err
	}
	return platform, err
}

func (r *platformRepository) CreatePlatform(ctx context.Context, platform entities.Platform) (entities.Platform, error) {
	err := r.db.Create(&platform).Error
	if err != nil {
		return entities.Platform{}, err
	}
	return platform, err
}

func (r *platformRepository) UpdatePlatform(ctx context.Context, platform entities.Platform) (entities.Platform, error) {
	err := r.db.Save(&platform).Error
	if err != nil {
		return entities.Platform{}, err
	}
	return platform, err
}

func (r *platformRepository) DeletePlatform(ctx context.Context, platformId int) error {
	err := r.db.Delete(&entities.Platform{}, "id = ?", platformId).Error
	if err != nil {
		return err
	}
	return nil
}
