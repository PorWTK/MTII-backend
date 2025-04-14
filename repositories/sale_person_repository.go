package repositories

import (
	"context"
	"mtii-backend/entities"

	"gorm.io/gorm"
)

type SalePersonRepository interface {
	GetAllSalePerson(ctx context.Context) ([]entities.SalePerson, error)
	GetSalePersonById(ctx context.Context, salePersonId int) (entities.SalePerson, error)
	CreateSalePerson(ctx context.Context, salePerson entities.SalePerson) (entities.SalePerson, error)
	UpdateSalePerson(ctx context.Context, salePerson entities.SalePerson) (entities.SalePerson, error)
	DeleteSalePerson(ctx context.Context, salePersonId int) error
}

type salePersonRepository struct {
	db *gorm.DB
}

func NewSalePersonRepository(db *gorm.DB) SalePersonRepository {
	return &salePersonRepository{
		db: db,
	}
}

func (r *salePersonRepository) GetAllSalePerson(ctx context.Context) ([]entities.SalePerson, error) {
	var salePeople []entities.SalePerson
	err := r.db.Find(&salePeople).Error
	if err != nil {
		return []entities.SalePerson{}, err
	}
	return salePeople, err
}

func (r *salePersonRepository) GetSalePersonById(ctx context.Context, salePersonId int) (entities.SalePerson, error) {
	var salePerson entities.SalePerson
	err := r.db.Where("id = ?", salePersonId).First(&salePerson).Error
	if err != nil {
		return entities.SalePerson{}, err
	}
	return salePerson, err
}

func (r *salePersonRepository) CreateSalePerson(ctx context.Context, salePerson entities.SalePerson) (entities.SalePerson, error) {
	err := r.db.Create(&salePerson).Error
	if err != nil {
		return entities.SalePerson{}, err
	}
	return salePerson, err
}

func (r *salePersonRepository) UpdateSalePerson(ctx context.Context, salePerson entities.SalePerson) (entities.SalePerson, error) {
	err := r.db.Save(&salePerson).Error
	if err != nil {
		return entities.SalePerson{}, err
	}
	return salePerson, err
}

func (r *salePersonRepository) DeleteSalePerson(ctx context.Context, salePersonId int) error {
	err := r.db.Delete(&entities.SalePerson{}, "id = ?", salePersonId).Error
	if err != nil {
		return err
	}
	return nil
}
