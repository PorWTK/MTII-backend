package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type ReceiverRepository interface {
	GetAllReceiver(ctx context.Context) ([]entities.Receiver, error)
	GetReceiverById(ctx context.Context, receiverId int) (entities.Receiver, error)
	CreateReceiver(ctx context.Context, receiver entities.Receiver) (entities.Receiver, error)
	UpdateReceiver(ctx context.Context, receiver entities.Receiver) (entities.Receiver, error)
	DeleteReceiver(ctx context.Context, receiverId int) error
}

type receiverRepository struct {
	db *gorm.DB
}

func NewReceiverRepository(db *gorm.DB) ReceiverRepository {
	return &receiverRepository{
		db: db,
	}
}

func (r *receiverRepository) GetAllReceiver(ctx context.Context) ([]entities.Receiver, error) {
	var receivers []entities.Receiver
	err := r.db.Find(&receivers).Error
	if err != nil {
		return []entities.Receiver{}, err
	}
	return receivers, err
}

func (r *receiverRepository) GetReceiverById(ctx context.Context, receiverId int) (entities.Receiver, error) {
	var receiver entities.Receiver
	err := r.db.Where("id = ?", receiverId).First(&receiver).Error
	if err != nil {
		return entities.Receiver{}, err
	}
	return receiver, err
}

func (r *receiverRepository) CreateReceiver(ctx context.Context, receiver entities.Receiver) (entities.Receiver, error) {
	err := r.db.Create(&receiver).Error
	if err != nil {
		return entities.Receiver{}, err
	}
	return receiver, err
}

func (r *receiverRepository) UpdateReceiver(ctx context.Context, receiver entities.Receiver) (entities.Receiver, error) {
	err := r.db.Save(&receiver).Error
	if err != nil {
		return entities.Receiver{}, err
	}
	return receiver, err
}

func (r *receiverRepository) DeleteReceiver(ctx context.Context, receiverId int) error {
	err := r.db.Delete(&entities.Receiver{}, "id = ?", receiverId).Error
	if err != nil {
		return err
	}
	return nil
}
