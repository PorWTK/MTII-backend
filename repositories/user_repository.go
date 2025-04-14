package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(ctx context.Context, userId int) (entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserById(ctx context.Context, userId int) (entities.User, error) {
	var user entities.User
	err := r.db.Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).Take(&user).Error
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
