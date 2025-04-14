package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type StatusRepository interface {
	GetAllStatus(ctx context.Context) ([]entities.Status, error)
	GetStatusById(ctx context.Context, statusId int) (entities.Status, error)
	CreateStatus(ctx context.Context, status entities.Status) (entities.Status, error)
	UpdateStatus(ctx context.Context, status entities.Status) (entities.Status, error)
	DeleteStatus(ctx context.Context, statusId int) error
}

type statusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) StatusRepository {
	return &statusRepository{
		db: db,
	}
}

func (r *statusRepository) GetAllStatus(ctx context.Context) ([]entities.Status, error) {
	var statuses []entities.Status
	err := r.db.Find(&statuses).Error
	if err != nil {
		return []entities.Status{}, err
	}
	return statuses, err
}

func (r *statusRepository) GetStatusById(ctx context.Context, statusId int) (entities.Status, error) {
	var status entities.Status
	err := r.db.Where("id = ?", statusId).First(&status).Error
	if err != nil {
		return entities.Status{}, err
	}
	return status, err
}

func (r *statusRepository) CreateStatus(ctx context.Context, status entities.Status) (entities.Status, error) {
	err := r.db.Create(&status).Error
	if err != nil {
		return entities.Status{}, err
	}
	return status, err
}

func (r *statusRepository) UpdateStatus(ctx context.Context, status entities.Status) (entities.Status, error) {
	err := r.db.Save(&status).Error
	if err != nil {
		return entities.Status{}, err
	}
	return status, err
}

func (r *statusRepository) DeleteStatus(ctx context.Context, statusId int) error {
	err := r.db.Delete(&entities.Status{}, "id = ?", statusId).Error
	if err != nil {
		return err
	}
	return nil
}
