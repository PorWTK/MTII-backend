package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/repositories"
)

type SalePersonService interface {
	GetAllSalePerson(ctx context.Context) ([]dtos.SalePerson, error)
	GetSalePersonById(ctx context.Context, salePersonId int) (dtos.SalePerson, error)
	CreateSalePerson(ctx context.Context, req dtos.SalePersonRequest) (dtos.SalePersonResponse, error)
	UpdateSalePerson(ctx context.Context, salePersonId int, req dtos.SalePersonRequest) (dtos.SalePersonResponse, error)
	DeleteSalePerson(ctx context.Context, salePersonId int) error
}

type salePersonService struct {
	salePersonRepository repositories.SalePersonRepository
}

func NewSalePersonService(
	salePersonRepository repositories.SalePersonRepository,
) SalePersonService {
	return &salePersonService{
		salePersonRepository: salePersonRepository,
	}
}

func (s *salePersonService) GetAllSalePerson(ctx context.Context) ([]dtos.SalePerson, error) {
	salePeople, err := s.salePersonRepository.GetAllSalePerson(ctx)
	if err != nil {
		return []dtos.SalePerson{}, fmt.Errorf("failed to get sale person: %w", err)
	}

	var salePersonDTOs []dtos.SalePerson
	for _, s := range salePeople {
		salePersonDTOs = append(salePersonDTOs, dtos.SalePerson{
			Id:   s.Id,
			Name: s.Name,
		})
	}

	if len(salePersonDTOs) == 0 {
		return []dtos.SalePerson{}, nil
	}

	return salePersonDTOs, nil
}

func (s *salePersonService) GetSalePersonById(ctx context.Context, salePersonId int) (dtos.SalePerson, error) {
	salePerson, err := s.salePersonRepository.GetSalePersonById(ctx, salePersonId)
	if err != nil {
		return dtos.SalePerson{}, fmt.Errorf("failed to get sale method: %w", err)
	}

	return dtos.SalePerson{
		Id:   salePerson.Id,
		Name: salePerson.Name,
	}, nil
}

func (s *salePersonService) CreateSalePerson(ctx context.Context, req dtos.SalePersonRequest) (dtos.SalePersonResponse, error) {
	data := entities.SalePerson{
		Name: req.Name,
	}

	salePerson, err := s.salePersonRepository.CreateSalePerson(ctx, data)
	if err != nil {
		return dtos.SalePersonResponse{}, fmt.Errorf("failed to save sale person: %w", err)
	}

	return dtos.SalePersonResponse{
		Id: salePerson.Id,
	}, nil
}

func (s *salePersonService) UpdateSalePerson(ctx context.Context, salePersonId int, req dtos.SalePersonRequest) (dtos.SalePersonResponse, error) {
	_, err := s.salePersonRepository.GetSalePersonById(ctx, salePersonId)
	if err != nil {
		return dtos.SalePersonResponse{}, fmt.Errorf("failed to get sale person: %w", err)
	}

	data := entities.SalePerson{
		Id:   salePersonId,
		Name: req.Name,
	}

	salePerson, err := s.salePersonRepository.UpdateSalePerson(ctx, data)
	if err != nil {
		return dtos.SalePersonResponse{}, fmt.Errorf("failed to save sale person: %w", err)
	}

	return dtos.SalePersonResponse{
		Id: salePerson.Id,
	}, nil
}

func (s *salePersonService) DeleteSalePerson(ctx context.Context, salePersonId int) error {
	salePerson, err := s.salePersonRepository.GetSalePersonById(ctx, salePersonId)
	if err != nil {
		return fmt.Errorf("failed to get sale person: %w", err)
	}

	err = s.salePersonRepository.DeleteSalePerson(ctx, salePerson.Id)
	if err != nil {
		return fmt.Errorf("failed to delete sale person: %w", err)
	}

	return nil
}
