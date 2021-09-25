package repository

import (
	"database/sql"

	"github.com/azcov/sagara_crud/cmd/product/app/domain"
)

type productRepository struct {
	*Queries
	db *sql.DB
}

// NewRepository creates a new store
func NewRepository(db *sql.DB) domain.ProductRepository {
	return &productRepository{
		db:      db,
		Queries: New(db),
	}
}
