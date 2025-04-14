package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/entities"
	"mtii-backend/repositories"
)

type ChannelService interface {
	GetAllChannel(ctx context.Context) ([]dtos.Channel, error)
	GetChannelById(ctx context.Context, channelId int) (dtos.Channel, error)
	CreateChannel(ctx context.Context, req dtos.ChannelRequest) (dtos.ChannelResponse, error)
	UpdateChannel(ctx context.Context, channelId int, req dtos.ChannelRequest) (dtos.ChannelResponse, error)
	DeleteChannel(ctx context.Context, channelId int) error
}

type channelService struct {
	channelRepository repositories.ChannelRepository
}

func NewChannelService(
	channelRepository repositories.ChannelRepository,
) ChannelService {
	return &channelService{
		channelRepository: channelRepository,
	}
}

func (s *channelService) GetAllChannel(ctx context.Context) ([]dtos.Channel, error) {
	channels, err := s.channelRepository.GetAllChannel(ctx)
	if err != nil {
		return []dtos.Channel{}, fmt.Errorf("failed to get channel: %w", err)
	}

	var channelDTOs []dtos.Channel
	for _, c := range channels {
		channelDTOs = append(channelDTOs, dtos.Channel{
			Id:   c.Id,
			Name: c.Name,
		})
	}

	if len(channelDTOs) == 0 {
		return []dtos.Channel{}, nil
	}

	return channelDTOs, nil
}

func (s *channelService) GetChannelById(ctx context.Context, channelId int) (dtos.Channel, error) {
	channel, err := s.channelRepository.GetChannelById(ctx, channelId)
	if err != nil {
		return dtos.Channel{}, fmt.Errorf("failed to get channel: %w", err)
	}

	return dtos.Channel{
		Id:   channel.Id,
		Name: channel.Name,
	}, nil
}

func (s *channelService) CreateChannel(ctx context.Context, req dtos.ChannelRequest) (dtos.ChannelResponse, error) {
	data := entities.Channel{
		Name: req.Name,
	}

	channel, err := s.channelRepository.CreateChannel(ctx, data)
	if err != nil {
		return dtos.ChannelResponse{}, fmt.Errorf("failed to save channel: %w", err)
	}

	return dtos.ChannelResponse{
		Id: channel.Id,
	}, nil
}

func (s *channelService) UpdateChannel(ctx context.Context, channelId int, req dtos.ChannelRequest) (dtos.ChannelResponse, error) {
	_, err := s.channelRepository.GetChannelById(ctx, channelId)
	if err != nil {
		return dtos.ChannelResponse{}, fmt.Errorf("failed to get channel: %w", err)
	}

	data := entities.Channel{
		Id:   channelId,
		Name: req.Name,
	}

	channel, err := s.channelRepository.UpdateChannel(ctx, data)
	if err != nil {
		return dtos.ChannelResponse{}, fmt.Errorf("failed to save channel: %w", err)
	}

	return dtos.ChannelResponse{
		Id: channel.Id,
	}, nil
}

func (s *channelService) DeleteChannel(ctx context.Context, channelId int) error {
	channel, err := s.channelRepository.GetChannelById(ctx, channelId)
	if err != nil {
		return fmt.Errorf("failed to get channel: %w", err)
	}

	err = s.channelRepository.DeleteChannel(ctx, channel.Id)
	if err != nil {
		return fmt.Errorf("failed to delete channel: %w", err)
	}

	return nil
}
