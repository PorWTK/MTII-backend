package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	GetAllPaymentMethod(ctx context.Context) ([]entities.PaymentMethod, error)
	GetPaymentMethodById(ctx context.Context, paymentMethodId int) (entities.PaymentMethod, error)
	CreatePaymentMethod(ctx context.Context, paymentMethod entities.PaymentMethod) (entities.PaymentMethod, error)
	UpdatePaymentMethod(ctx context.Context, paymentMethod entities.PaymentMethod) (entities.PaymentMethod, error)
	DeletePaymentMethod(ctx context.Context, paymentMethodId int) error
}

type paymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}

func (r *paymentMethodRepository) GetAllPaymentMethod(ctx context.Context) ([]entities.PaymentMethod, error) {
	var paymentMethods []entities.PaymentMethod
	err := r.db.Find(&paymentMethods).Error
	if err != nil {
		return []entities.PaymentMethod{}, err
	}
	return paymentMethods, err
}

func (r *paymentMethodRepository) GetPaymentMethodById(ctx context.Context, paymentMethodId int) (entities.PaymentMethod, error) {
	var paymentMethod entities.PaymentMethod
	err := r.db.Where("id = ?", paymentMethodId).First(&paymentMethod).Error
	if err != nil {
		return entities.PaymentMethod{}, err
	}
	return paymentMethod, err
}

func (r *paymentMethodRepository) CreatePaymentMethod(ctx context.Context, paymentMethod entities.PaymentMethod) (entities.PaymentMethod, error) {
	err := r.db.Create(&paymentMethod).Error
	if err != nil {
		return entities.PaymentMethod{}, err
	}
	return paymentMethod, err
}

func (r *paymentMethodRepository) UpdatePaymentMethod(ctx context.Context, paymentMethod entities.PaymentMethod) (entities.PaymentMethod, error) {
	err := r.db.Save(&paymentMethod).Error
	if err != nil {
		return entities.PaymentMethod{}, err
	}
	return paymentMethod, err
}

func (r *paymentMethodRepository) DeletePaymentMethod(ctx context.Context, paymentMethodId int) error {
	err := r.db.Delete(&entities.PaymentMethod{}, "id = ?", paymentMethodId).Error
	if err != nil {
		return err
	}
	return nil
}
