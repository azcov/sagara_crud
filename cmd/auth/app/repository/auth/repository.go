package repository

import (
	"database/sql"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
)

type sessionRepository struct {
	*Queries
	db *sql.DB
}

// NewRepository creates a new store
func NewRepository(db *sql.DB) domain.AuthenticationRepository {
	return &sessionRepository{
		db:      db,
		Queries: New(db),
	}
}
