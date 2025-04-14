package services

import (
	"context"
	"fmt"
	"mtii-backend/dtos"
	"mtii-backend/helpers"
	"mtii-backend/repositories"
	"time"
)

type UserService interface {
	VerifyCredential(ctx context.Context, req dtos.LoginRequest) (dtos.LoginResponse, error)
}

type userService struct {
	tokenService   TokenService
	userRepository repositories.UserRepository
}

func NewUserService(
	tokenService TokenService,
	userRepository repositories.UserRepository,
) UserService {
	return &userService{
		tokenService:   tokenService,
		userRepository: userRepository,
	}
}

func (s *userService) VerifyCredential(ctx context.Context, req dtos.LoginRequest) (dtos.LoginResponse, error) {
	user, err := s.userRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return dtos.LoginResponse{}, fmt.Errorf("username not registered")
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dtos.LoginResponse{}, fmt.Errorf("wrong password")
	}

	token := s.tokenService.GenerateToken(user.Id)
	return dtos.LoginResponse{
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}, nil
}
