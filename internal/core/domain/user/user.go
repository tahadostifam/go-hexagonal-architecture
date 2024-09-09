package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Username Username
	Age      int
}

func New(name string, username Username, age int) User {
	return User{
		ID:       uuid.New(),
		Name:     name,
		Username: username,
		Age:      age,
	}
}
