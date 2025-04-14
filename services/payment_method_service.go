package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/repositories"
)

type PaymentMethodService interface {
	GetAllPaymentMethod(ctx context.Context) ([]dtos.PaymentMethod, error)
	GetPaymentMethodById(ctx context.Context, paymentMethodId int) (dtos.PaymentMethod, error)
	CreatePaymentMethod(ctx context.Context, req dtos.PaymentMethodRequest) (dtos.PaymentMethodResponse, error)
	UpdatePaymentMethod(ctx context.Context, paymentMethodId int, req dtos.PaymentMethodRequest) (dtos.PaymentMethodResponse, error)
	DeletePaymentMethod(ctx context.Context, paymentMethodId int) error
}

type paymentMethodService struct {
	paymentMethodRepository repositories.PaymentMethodRepository
}

func NewPaymentMethodService(
	paymentMethodRepository repositories.PaymentMethodRepository,
) PaymentMethodService {
	return &paymentMethodService{
		paymentMethodRepository: paymentMethodRepository,
	}
}

func (s *paymentMethodService) GetAllPaymentMethod(ctx context.Context) ([]dtos.PaymentMethod, error) {
	paymentMethods, err := s.paymentMethodRepository.GetAllPaymentMethod(ctx)
	if err != nil {
		return []dtos.PaymentMethod{}, fmt.Errorf("failed to get payment method: %w", err)
	}

	var paymentMethodDTOs []dtos.PaymentMethod
	for _, p := range paymentMethods {
		paymentMethodDTOs = append(paymentMethodDTOs, dtos.PaymentMethod{
			Id:   p.Id,
			Name: p.Name,
		})
	}

	if len(paymentMethodDTOs) == 0 {
		return []dtos.PaymentMethod{}, nil
	}

	return paymentMethodDTOs, nil
}

func (s *paymentMethodService) GetPaymentMethodById(ctx context.Context, paymentMethodId int) (dtos.PaymentMethod, error) {
	paymentMethod, err := s.paymentMethodRepository.GetPaymentMethodById(ctx, paymentMethodId)
	if err != nil {
		return dtos.PaymentMethod{}, fmt.Errorf("failed to get payment method: %w", err)
	}

	return dtos.PaymentMethod{
		Id:   paymentMethod.Id,
		Name: paymentMethod.Name,
	}, nil
}

func (s *paymentMethodService) CreatePaymentMethod(ctx context.Context, req dtos.PaymentMethodRequest) (dtos.PaymentMethodResponse, error) {
	data := entities.PaymentMethod{
		Name: req.Name,
	}

	paymentMethod, err := s.paymentMethodRepository.CreatePaymentMethod(ctx, data)
	if err != nil {
		return dtos.PaymentMethodResponse{}, fmt.Errorf("failed to save payment method: %w", err)
	}

	return dtos.PaymentMethodResponse{
		Id: paymentMethod.Id,
	}, nil
}

func (s *paymentMethodService) UpdatePaymentMethod(ctx context.Context, paymentMethodId int, req dtos.PaymentMethodRequest) (dtos.PaymentMethodResponse, error) {
	_, err := s.paymentMethodRepository.GetPaymentMethodById(ctx, paymentMethodId)
	if err != nil {
		return dtos.PaymentMethodResponse{}, fmt.Errorf("failed to get payment method: %w", err)
	}

	data := entities.PaymentMethod{
		Id:   paymentMethodId,
		Name: req.Name,
	}

	paymentMethod, err := s.paymentMethodRepository.UpdatePaymentMethod(ctx, data)
	if err != nil {
		return dtos.PaymentMethodResponse{}, fmt.Errorf("failed to save payment method: %w", err)
	}

	return dtos.PaymentMethodResponse{
		Id: paymentMethod.Id,
	}, nil
}

func (s *paymentMethodService) DeletePaymentMethod(ctx context.Context, paymentMethodId int) error {
	paymentMethod, err := s.paymentMethodRepository.GetPaymentMethodById(ctx, paymentMethodId)
	if err != nil {
		return fmt.Errorf("failed to get payment method: %w", err)
	}

	err = s.paymentMethodRepository.DeletePaymentMethod(ctx, paymentMethod.Id)
	if err != nil {
		return fmt.Errorf("failed to delete payment method: %w", err)
	}

	return nil
}
