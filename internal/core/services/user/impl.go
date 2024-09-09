package user_service

import (
	"context"
	"fmt"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/ports"
)

type Service struct {
	userRepo ports.UserRepositorySecondaryPort
}

func NewService(userRepo ports.UserRepositorySecondaryPort) *Service {
	return &Service{userRepo}
}

func (s *Service) CreateAccount(ctx context.Context, name, usernameStr string, age int) (*user.User, error) {
	username, err := user.NewUsername(usernameStr)
	if err != nil {
		return nil, err
	}

	user := user.New(name, username, age)

	if err := s.userRepo.Add(ctx, &user); err != nil {
		return nil, fmt.Errorf("failed to add the user: %w", err)
	}

	return &user, nil
}
