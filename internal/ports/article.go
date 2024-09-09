package ports

import (
	"context"
	"errors"

	"github.com/samverrall/hex-structure/internal/core/domain/article"
)

var (
	ErrArticleNotFound = errors.New("article does not exist")
)

type ArticleRepository interface {
	Add(context.Context, article.Article) (*article.Article, error)
}
