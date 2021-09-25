package repository

import (
	"database/sql"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
)

type authRepository struct {
	*Queries
	db *sql.DB
}

// NewRepository creates a new store
func NewRepository(db *sql.DB) domain.SessionRepository {
	return &authRepository{
		db:      db,
		Queries: New(db),
	}
}
