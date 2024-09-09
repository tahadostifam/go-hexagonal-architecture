package ports

import (
	"context"
	"errors"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type UserRepositorySecondaryPort interface {
	Add(context.Context, *user.User) error
	Update(context.Context, user.Username, map[string]interface{}) error
}
