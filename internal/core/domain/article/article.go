package article

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID        uuid.UUID
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title, content string, createdAt time.Time) Article {
	return Article{
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
	}
}
