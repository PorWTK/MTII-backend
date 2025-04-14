package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/repositories"
)

type BankService interface {
	GetAllBank(ctx context.Context) ([]dtos.Bank, error)
	GetBankById(ctx context.Context, bankId int) (dtos.Bank, error)
	CreateBank(ctx context.Context, req dtos.BankRequest) (dtos.BankResponse, error)
	UpdateBank(ctx context.Context, bankId int, req dtos.BankRequest) (dtos.BankResponse, error)
	DeleteBank(ctx context.Context, bankId int) error
}

type bankService struct {
	bankRepository repositories.BankRepository
}

func NewBankService(
	bankRepository repositories.BankRepository,
) BankService {
	return &bankService{
		bankRepository: bankRepository,
	}
}

func (s *bankService) GetAllBank(ctx context.Context) ([]dtos.Bank, error) {
	banks, err := s.bankRepository.GetAllBank(ctx)
	if err != nil {
		return []dtos.Bank{}, fmt.Errorf("failed to get the bank: %w", err)
	}

	var bankDTOs []dtos.Bank
	for _, p := range banks {
		bankDTOs = append(bankDTOs, dtos.Bank{
			Id:   p.Id,
			Name: p.Name,
		})
	}

	if len(bankDTOs) == 0 {
		return []dtos.Bank{}, nil
	}

	return bankDTOs, nil
}

func (s *bankService) GetBankById(ctx context.Context, bankId int) (dtos.Bank, error) {
	bank, err := s.bankRepository.GetBankById(ctx, bankId)
	if err != nil {
		return dtos.Bank{}, fmt.Errorf("failed to get the bank: %w", err)
	}

	return dtos.Bank{
		Id:   bank.Id,
		Name: bank.Name,
	}, nil
}

func (s *bankService) CreateBank(ctx context.Context, req dtos.BankRequest) (dtos.BankResponse, error) {
	data := entities.Bank{
		Name: req.Name,
	}

	bank, err := s.bankRepository.CreateBank(ctx, data)
	if err != nil {
		return dtos.BankResponse{}, fmt.Errorf("failed to save the bank: %w", err)
	}

	return dtos.BankResponse{
		Id: bank.Id,
	}, nil
}

func (s *bankService) UpdateBank(ctx context.Context, bankId int, req dtos.BankRequest) (dtos.BankResponse, error) {
	_, err := s.bankRepository.GetBankById(ctx, bankId)
	if err != nil {
		return dtos.BankResponse{}, fmt.Errorf("failed to get the bank: %w", err)
	}

	data := entities.Bank{
		Id:   bankId,
		Name: req.Name,
	}

	bank, err := s.bankRepository.UpdateBank(ctx, data)
	if err != nil {
		return dtos.BankResponse{}, fmt.Errorf("failed to save the bank: %w", err)
	}

	return dtos.BankResponse{
		Id: bank.Id,
	}, nil
}

func (s *bankService) DeleteBank(ctx context.Context, bankId int) error {
	bank, err := s.bankRepository.GetBankById(ctx, bankId)
	if err != nil {
		return fmt.Errorf("failed to get the bank: %w", err)
	}

	err = s.bankRepository.DeleteBank(ctx, bank.Id)
	if err != nil {
		return fmt.Errorf("failed to delete the bank: %w", err)
	}

	return nil
}
