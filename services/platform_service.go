package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/repositories"
)

type PlatformService interface {
	GetAllPlatform(ctx context.Context) ([]dtos.Platform, error)
	GetPlatformById(ctx context.Context, platformId int) (dtos.Platform, error)
	CreatePlatform(ctx context.Context, req dtos.PlatformRequest) (dtos.PlatformResponse, error)
	UpdatePlatform(ctx context.Context, platformId int, req dtos.PlatformRequest) (dtos.PlatformResponse, error)
	DeletePlatform(ctx context.Context, platformId int) error
}

type platformService struct {
	platformRepository repositories.PlatformRepository
}

func NewPlatformService(
	platformRepository repositories.PlatformRepository,
) PlatformService {
	return &platformService{
		platformRepository: platformRepository,
	}
}

func (s *platformService) GetAllPlatform(ctx context.Context) ([]dtos.Platform, error) {
	platforms, err := s.platformRepository.GetAllPlatform(ctx)
	if err != nil {
		return []dtos.Platform{}, fmt.Errorf("failed to get platform: %w", err)
	}

	var platformDTOs []dtos.Platform
	for _, p := range platforms {
		platformDTOs = append(platformDTOs, dtos.Platform{
			Id:   p.Id,
			Name: p.Name,
		})
	}

	if len(platformDTOs) == 0 {
		return []dtos.Platform{}, nil
	}

	return platformDTOs, nil
}

func (s *platformService) GetPlatformById(ctx context.Context, platformId int) (dtos.Platform, error) {
	platform, err := s.platformRepository.GetPlatformById(ctx, platformId)
	if err != nil {
		return dtos.Platform{}, fmt.Errorf("failed to get platform: %w", err)
	}

	return dtos.Platform{
		Id:   platform.Id,
		Name: platform.Name,
	}, nil
}

func (s *platformService) CreatePlatform(ctx context.Context, req dtos.PlatformRequest) (dtos.PlatformResponse, error) {
	data := entities.Platform{
		Name: req.Name,
	}

	platform, err := s.platformRepository.CreatePlatform(ctx, data)
	if err != nil {
		return dtos.PlatformResponse{}, fmt.Errorf("failed to save platform: %w", err)
	}

	return dtos.PlatformResponse{
		Id: platform.Id,
	}, nil
}

func (s *platformService) UpdatePlatform(ctx context.Context, platformId int, req dtos.PlatformRequest) (dtos.PlatformResponse, error) {
	_, err := s.platformRepository.GetPlatformById(ctx, platformId)
	if err != nil {
		return dtos.PlatformResponse{}, fmt.Errorf("failed to get platform: %w", err)
	}

	data := entities.Platform{
		Id:   platformId,
		Name: req.Name,
	}

	platform, err := s.platformRepository.UpdatePlatform(ctx, data)
	if err != nil {
		return dtos.PlatformResponse{}, fmt.Errorf("failed to save platform: %w", err)
	}

	return dtos.PlatformResponse{
		Id: platform.Id,
	}, nil
}

func (s *platformService) DeletePlatform(ctx context.Context, platformId int) error {
	platform, err := s.platformRepository.GetPlatformById(ctx, platformId)
	if err != nil {
		return fmt.Errorf("failed to get platform: %w", err)
	}

	err = s.platformRepository.DeletePlatform(ctx, platform.Id)
	if err != nil {
		return fmt.Errorf("failed to delete platform: %w", err)
	}

	return nil
}
