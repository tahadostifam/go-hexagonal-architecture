package sqlite_adapter

import (
	"context"

	"github.com/samverrall/hex-structure/internal/core/domain/user"
	"github.com/samverrall/hex-structure/internal/ports"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(dialector gorm.Dialector) (ports.UserRepository, error) {
	db, err := GORM(dialector)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return &UserRepository{db}, nil
}

func (r *UserRepository) Add(ctx context.Context, user *user.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) Update(ctx context.Context, username user.Username, changes map[string]interface{}) error {
	// r.db.Model(&user.User{}).Where("username = ?", username)
	panic("unimplemented")
}
