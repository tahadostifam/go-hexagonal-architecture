package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsername(t *testing.T) {
	testCases := []struct {
		Username string
		Valid    bool
	}{
		{
			Username: "john",
			Valid:    true,
		},
		{
			Username: "john_doe",
			Valid:    true,
		},
		{
			Username: "johndoe",
			Valid:    true,
		},
		{
			Username: "Johndoe",
			Valid:    false,
		},
		{
			Username: "John Doe",
			Valid:    false,
		},
		{
			Username: "john doe",
			Valid:    false,
		},
	}

	for _, tc := range testCases {
		_, err := NewUsername(tc.Username)
		assert.Equal(t, (err == nil), tc.Valid, "invalid violation")
	}
}
