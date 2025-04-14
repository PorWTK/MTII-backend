package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/helpers"
	"mtii-backend/repositories"
)

type ReceiverService interface {
	GetAllReceiver(ctx context.Context) ([]dtos.Receiver, error)
	GetReceiverById(ctx context.Context, receiverId int) (dtos.Receiver, error)
	CreateReceiver(ctx context.Context, req dtos.CreateReceiverRequest) (dtos.ReceiverResponse, error)
	UpdateReceiver(ctx context.Context, receiverId int, req dtos.UpdateReceiverRequest) (dtos.ReceiverResponse, error)
	DeleteReceiver(ctx context.Context, receiverId int) error
}

type receiverService struct {
	receiverRepository repositories.ReceiverRepository
}

func NewReceiverService(
	receiverRepository repositories.ReceiverRepository,
) ReceiverService {
	return &receiverService{
		receiverRepository: receiverRepository,
	}
}

func (s *receiverService) GetAllReceiver(ctx context.Context) ([]dtos.Receiver, error) {
	receivers, err := s.receiverRepository.GetAllReceiver(ctx)
	if err != nil {
		return []dtos.Receiver{}, fmt.Errorf("failed to get receiver: %w", err)
	}

	var receiverDTOs []dtos.Receiver
	for _, r := range receivers {
		receiverDTOs = append(receiverDTOs, dtos.Receiver{
			Id:         r.Id,
			Name:       r.Name,
			Address:    r.Address,
			Email:      r.Email,
			Phone:      r.Phone,
			TaxPayerId: r.TaxPayerId,
		})
	}

	if len(receiverDTOs) == 0 {
		return []dtos.Receiver{}, nil
	}

	return receiverDTOs, nil
}

func (s *receiverService) GetReceiverById(ctx context.Context, receiverId int) (dtos.Receiver, error) {
	receiver, err := s.receiverRepository.GetReceiverById(ctx, receiverId)
	if err != nil {
		return dtos.Receiver{}, fmt.Errorf("failed to get receiver: %w", err)
	}

	return dtos.Receiver{
		Id:         receiver.Id,
		Name:       receiver.Name,
		Address:    receiver.Address,
		Email:      receiver.Email,
		Phone:      receiver.Phone,
		TaxPayerId: receiver.TaxPayerId,
	}, nil
}

func (s *receiverService) CreateReceiver(ctx context.Context, req dtos.CreateReceiverRequest) (dtos.ReceiverResponse, error) {
	data := entities.Receiver{
		Name:       req.Name,
		Address:    req.Address,
		Email:      req.Email,
		Phone:      req.Phone,
		TaxPayerId: req.TaxPayerId,
	}

	receiver, err := s.receiverRepository.CreateReceiver(ctx, data)
	if err != nil {
		return dtos.ReceiverResponse{}, fmt.Errorf("failed to save receiver: %w", err)
	}

	return dtos.ReceiverResponse{
		Id: receiver.Id,
	}, nil
}

func (s *receiverService) UpdateReceiver(ctx context.Context, receiverId int, req dtos.UpdateReceiverRequest) (dtos.ReceiverResponse, error) {
	receiver, err := s.receiverRepository.GetReceiverById(ctx, receiverId)
	if err != nil {
		return dtos.ReceiverResponse{}, fmt.Errorf("failed to get receiver: %w", err)
	}

	data := entities.Receiver{
		Id:         receiverId,
		Name:       helpers.DefaultIfEmpty(req.Name, receiver.Name),
		Address:    helpers.DefaultIfEmpty(req.Address, receiver.Address),
		Email:      helpers.DefaultIfEmpty(req.Email, receiver.Email),
		Phone:      helpers.DefaultIfEmpty(req.Phone, receiver.Phone),
		TaxPayerId: helpers.DefaultIfEmpty(req.TaxPayerId, receiver.TaxPayerId),
	}

	updatedReceiver, err := s.receiverRepository.UpdateReceiver(ctx, data)
	if err != nil {
		return dtos.ReceiverResponse{}, fmt.Errorf("failed to save receiver: %w", err)
	}

	return dtos.ReceiverResponse{
		Id: updatedReceiver.Id,
	}, nil
}

func (s *receiverService) DeleteReceiver(ctx context.Context, receiverId int) error {
	receiver, err := s.receiverRepository.GetReceiverById(ctx, receiverId)
	if err != nil {
		return fmt.Errorf("failed to get receiver: %w", err)
	}

	err = s.receiverRepository.DeleteReceiver(ctx, receiver.Id)
	if err != nil {
		return fmt.Errorf("failed to delete receiver: %w", err)
	}

	return nil
}
