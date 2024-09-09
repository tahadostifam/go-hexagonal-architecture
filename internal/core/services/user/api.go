package user_service

import (
	"context"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
)

type Api interface {
	CreateAccount(ctx context.Context, name, username string, age int) (*user.User, error)
}
