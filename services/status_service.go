package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/repositories"
)

type StatusService interface {
	GetAllStatus(ctx context.Context) ([]dtos.Status, error)
	GetStatusById(ctx context.Context, statusId int) (dtos.Status, error)
	CreateStatus(ctx context.Context, req dtos.StatusRequest) (dtos.StatusResponse, error)
	UpdateStatus(ctx context.Context, statusId int, req dtos.StatusRequest) (dtos.StatusResponse, error)
	DeleteStatus(ctx context.Context, statusId int) error
}

type statusService struct {
	statusRepository repositories.StatusRepository
}

func NewStatusService(
	statusRepository repositories.StatusRepository,
) StatusService {
	return &statusService{
		statusRepository: statusRepository,
	}
}

func (s *statusService) GetAllStatus(ctx context.Context) ([]dtos.Status, error) {
	statuses, err := s.statusRepository.GetAllStatus(ctx)
	if err != nil {
		return []dtos.Status{}, fmt.Errorf("failed to get status: %w", err)
	}

	var statusDTOs []dtos.Status
	for _, s := range statuses {
		statusDTOs = append(statusDTOs, dtos.Status{
			Id:   s.Id,
			Name: s.Name,
		})
	}

	if len(statusDTOs) == 0 {
		return []dtos.Status{}, nil
	}

	return statusDTOs, nil
}

func (s *statusService) GetStatusById(ctx context.Context, statusId int) (dtos.Status, error) {
	status, err := s.statusRepository.GetStatusById(ctx, statusId)
	if err != nil {
		return dtos.Status{}, fmt.Errorf("failed to get status: %w", err)
	}

	return dtos.Status{
		Id:   status.Id,
		Name: status.Name,
	}, nil
}

func (s *statusService) CreateStatus(ctx context.Context, req dtos.StatusRequest) (dtos.StatusResponse, error) {
	data := entities.Status{
		Name: req.Name,
	}

	status, err := s.statusRepository.CreateStatus(ctx, data)
	if err != nil {
		return dtos.StatusResponse{}, fmt.Errorf("failed to save status: %w", err)
	}

	return dtos.StatusResponse{
		Id: status.Id,
	}, nil
}

func (s *statusService) UpdateStatus(ctx context.Context, statusId int, req dtos.StatusRequest) (dtos.StatusResponse, error) {
	_, err := s.statusRepository.GetStatusById(ctx, statusId)
	if err != nil {
		return dtos.StatusResponse{}, fmt.Errorf("failed to get status: %w", err)
	}

	data := entities.Status{
		Id:   statusId,
		Name: req.Name,
	}

	status, err := s.statusRepository.UpdateStatus(ctx, data)
	if err != nil {
		return dtos.StatusResponse{}, fmt.Errorf("failed to save status: %w", err)
	}

	return dtos.StatusResponse{
		Id: status.Id,
	}, nil
}

func (s *statusService) DeleteStatus(ctx context.Context, statusId int) error {
	status, err := s.statusRepository.GetStatusById(ctx, statusId)
	if err != nil {
		return fmt.Errorf("failed to get status: %w", err)
	}

	err = s.statusRepository.DeleteStatus(ctx, status.Id)
	if err != nil {
		return fmt.Errorf("failed to delete status: %w", err)
	}

	return nil
}
