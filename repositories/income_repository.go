package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type IncomeRepository interface {
	GetAllIncome(ctx context.Context) ([]entities.Income, error)
	GetIncomeByInvoiceIdNumber(ctx context.Context, incomeInvoiceIdNumber int) (entities.Income, error)
	CreateIncome(ctx context.Context, income entities.Income) (entities.Income, error)
	UpdateIncome(ctx context.Context, income entities.Income) (entities.Income, error)
	UpdateIncomeWithNewInvoiceIdNumber(ctx context.Context, income entities.Income, oldInvoiceIdNumber int) (entities.Income, error)
	DeleteIncome(ctx context.Context, incomeInvoiceIdNumber int) error
}

type incomeRepository struct {
	db *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) IncomeRepository {
	return &incomeRepository{
		db: db,
	}
}

func (r *incomeRepository) GetAllIncome(ctx context.Context) ([]entities.Income, error) {
	var incomes []entities.Income
	err := r.db.
		Preload("Platform").
		Preload("Status").
		Preload("PaymentMethod").
		Preload("Receiver").
		Preload("SalePerson").
		Preload("Channel").
		Preload("Bank").
		Find(&incomes).Error
	if err != nil {
		return []entities.Income{}, err
	}
	return incomes, err
}

func (r *incomeRepository) GetIncomeByInvoiceIdNumber(ctx context.Context, incomeInvoiceIdNumber int) (entities.Income, error) {
	var income entities.Income
	err := r.db.
		Preload("Platform").
		Preload("Status").
		Preload("PaymentMethod").
		Preload("Receiver").
		Preload("SalePerson").
		Preload("Channel").
		Preload("Bank").
		Where("invoice_id_number = ?", incomeInvoiceIdNumber).
		First(&income).Error
	if err != nil {
		return entities.Income{}, err
	}
	return income, err
}

func (r *incomeRepository) CreateIncome(ctx context.Context, income entities.Income) (entities.Income, error) {
	err := r.db.Create(&income).Error
	if err != nil {
		return entities.Income{}, err
	}
	return income, err
}

func (r *incomeRepository) UpdateIncome(ctx context.Context, income entities.Income) (entities.Income, error) {
	err := r.db.Save(&income).Error
	if err != nil {
		return entities.Income{}, err
	}
	return income, err
}

func (r *incomeRepository) UpdateIncomeWithNewInvoiceIdNumber(ctx context.Context, income entities.Income, oldInvoiceIdNumber int) (entities.Income, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return entities.Income{}, tx.Error
	}

	if err := tx.Create(&income).Error; err != nil {
		tx.Rollback()
		return entities.Income{}, err
	}

	if err := tx.Model(&entities.Detail{}).Where("income_invoice_id_number = ?", oldInvoiceIdNumber).
		Update("income_invoice_id_number", income.InvoiceIdNumber).Error; err != nil {
		tx.Rollback()
		return entities.Income{}, err
	}

	if err := tx.Where("invoice_id_number = ?", oldInvoiceIdNumber).Delete(&entities.Income{}).Error; err != nil {
		tx.Rollback()
		return entities.Income{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.Income{}, err
	}

	return income, nil
}

func (r *incomeRepository) DeleteIncome(ctx context.Context, incomeInvoiceIdNumber int) error {
	err := r.db.Delete(&entities.Income{}, "invoice_id_number = ?", incomeInvoiceIdNumber).Error
	if err != nil {
		return err
	}
	return nil
}
