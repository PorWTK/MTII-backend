package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type BankRepository interface {
	GetAllBank(ctx context.Context) ([]entities.Bank, error)
	GetBankById(ctx context.Context, bankId int) (entities.Bank, error)
	CreateBank(ctx context.Context, bank entities.Bank) (entities.Bank, error)
	UpdateBank(ctx context.Context, bank entities.Bank) (entities.Bank, error)
	DeleteBank(ctx context.Context, bankId int) error
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{
		db: db,
	}
}

func (r *bankRepository) GetAllBank(ctx context.Context) ([]entities.Bank, error) {
	var banks []entities.Bank
	err := r.db.Find(&banks).Error
	if err != nil {
		return []entities.Bank{}, err
	}
	return banks, err
}

func (r *bankRepository) GetBankById(ctx context.Context, bankId int) (entities.Bank, error) {
	var bank entities.Bank
	err := r.db.Where("id = ?", bankId).First(&bank).Error
	if err != nil {
		return entities.Bank{}, err
	}
	return bank, err
}

func (r *bankRepository) CreateBank(ctx context.Context, bank entities.Bank) (entities.Bank, error) {
	err := r.db.Create(&bank).Error
	if err != nil {
		return entities.Bank{}, err
	}
	return bank, err
}

func (r *bankRepository) UpdateBank(ctx context.Context, bank entities.Bank) (entities.Bank, error) {
	err := r.db.Save(&bank).Error
	if err != nil {
		return entities.Bank{}, err
	}
	return bank, err
}

func (r *bankRepository) DeleteBank(ctx context.Context, bankId int) error {
	err := r.db.Delete(&entities.Bank{}, "id = ?", bankId).Error
	if err != nil {
		return err
	}
	return nil
}
