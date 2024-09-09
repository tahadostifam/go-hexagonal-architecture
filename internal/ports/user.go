package ports

import (
	"context"
	"errors"

	"github.com/samverrall/hex-structure/internal/core/domain/user"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
)

type UserRepository interface {
	Add(context.Context, *user.User) error
	Update(context.Context, user.Username, map[string]interface{}) error
}
