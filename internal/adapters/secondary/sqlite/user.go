package sqlite_adapter

import (
	"context"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/user"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/ports"
	"gorm.io/gorm"
)

type UserRepositorySecondaryPort struct {
	db *gorm.DB
}

func NewUserRepositorySecondaryPort(dialector gorm.Dialector) (ports.UserRepositorySecondaryPort, error) {
	db, err := GORM(dialector)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return &UserRepositorySecondaryPort{db}, nil
}

func (r *UserRepositorySecondaryPort) Add(ctx context.Context, user *user.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepositorySecondaryPort) Update(ctx context.Context, username user.Username, changes map[string]interface{}) error {
	panic("unimplemented")
}
