package user_service

import (
	"context"

	"github.com/samverrall/hex-structure/internal/core/domain/user"
)

type Api interface {
	CreateAccount(ctx context.Context, name, username string, age int) (*user.User, error)
}
