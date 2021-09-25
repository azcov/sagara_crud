package repository

import (
	"database/sql"

	"github.com/azcov/sagara_crud/cmd/article/app/domain"
)

type articleRepository struct {
	*Queries
	db *sql.DB
}

// NewRepository creates a new store
func NewRepository(db *sql.DB) domain.ArticleRepository {
	return &articleRepository{
		db:      db,
		Queries: New(db),
	}
}
