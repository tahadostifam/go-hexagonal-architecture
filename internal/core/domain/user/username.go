package user

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrEmptyUsername   = errors.New("username can't be empty")
	ErrInvalidUsername = errors.New("invalid username")
)

type Username string

const EmptyUsername Username = ""

func NewUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)

	if username == "" {
		return "", ErrEmptyUsername
	}

	// Validate username
	rxp := regexp.MustCompile(`^[a-z_]+$`)
	isValid := rxp.MatchString(username)

	if isValid {
		return Username(username), nil
	}

	return EmptyUsername, ErrInvalidUsername
}
