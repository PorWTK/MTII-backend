package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type DetailRepository interface {
	GetAllDetail(ctx context.Context) ([]entities.Detail, error)
	GetDetailById(ctx context.Context, detailId int) (entities.Detail, error)
	CreateDetail(ctx context.Context, detail entities.Detail) (entities.Detail, error)
	UpdateDetail(ctx context.Context, detail entities.Detail) (entities.Detail, error)
	DeleteDetail(ctx context.Context, detailId int) error
}

type detailRepository struct {
	db *gorm.DB
}

func NewDetailRepository(db *gorm.DB) DetailRepository {
	return &detailRepository{
		db: db,
	}
}

func (r *detailRepository) GetAllDetail(ctx context.Context) ([]entities.Detail, error) {
	var details []entities.Detail
	err := r.db.
		Preload("Income").
		Preload("Income.Platform").
		Preload("Income.Status").
		Preload("Income.PaymentMethod").
		Preload("Income.Receiver").
		Preload("Income.SalePerson").
		Preload("Income.Channel").
		Preload("Income.Bank").
		Find(&details).Error
	if err != nil {
		return []entities.Detail{}, err
	}
	return details, err
}

func (r *detailRepository) GetDetailById(ctx context.Context, detailId int) (entities.Detail, error) {
	var detail entities.Detail
	err := r.db.
		Preload("Income").
		Preload("Income.Platform").
		Preload("Income.Status").
		Preload("Income.PaymentMethod").
		Preload("Income.Receiver").
		Preload("Income.SalePerson").
		Preload("Income.Channel").
		Preload("Income.Bank").
		Where("id = ?", detailId).
		First(&detail).Error
	if err != nil {
		return entities.Detail{}, err
	}
	return detail, err
}

func (r *detailRepository) CreateDetail(ctx context.Context, detail entities.Detail) (entities.Detail, error) {
	err := r.db.Create(&detail).Error
	if err != nil {
		return entities.Detail{}, err
	}
	return detail, err
}

func (r *detailRepository) UpdateDetail(ctx context.Context, detail entities.Detail) (entities.Detail, error) {
	err := r.db.Save(&detail).Error
	if err != nil {
		return entities.Detail{}, err
	}
	return detail, err
}

func (r *detailRepository) DeleteDetail(ctx context.Context, detailId int) error {
	err := r.db.Delete(&entities.Detail{}, "id = ?", detailId).Error
	if err != nil {
		return err
	}
	return nil
}
