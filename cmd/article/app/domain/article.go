package domain

import (
	"database/sql"

	"github.com/google/uuid"
)

type CreateArticleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetArticleRequest struct {
	ArticleID uuid.UUID `param:"article_id" json:"-"`
}
type GetArticlesRequest struct {
	Search string `query:"search"`
	Page   int32  `query:"page"`
	Limit  int32  `query:"limit"`
	SortBy string `query:"sort_by"`
}

type UpdateArticleRequest struct {
	ArticleID   uuid.UUID `param:"article_id" json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type DeleteArticleRequest struct {
	ArticleID uuid.UUID `param:"article_id" json:"-"`
}

type ForceDeleteArticleRequest struct {
	ArticleID uuid.UUID `param:"article_id" json:"-"`
}

type GetArticleResponse struct {
	ArticleID   string  `json:"article_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CreatedAt   int64   `json:"created_at"`
	CreatedBy   string  `json:"created_by"`
	UpdatedAt   *int64  `json:"updated_at"`
	UpdatedBy   *string `json:"updated_by"`
}

type InsertArticleParams struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   int64     `json:"created_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
}

type DeleteArticleParams struct {
	DeletedAt sql.NullInt64 `json:"deleted_at"`
	DeletedBy uuid.NullUUID `json:"deleted_by"`
	ArticleID uuid.UUID     `json:"article_id"`
}

type GetArticleByArticleIDRow struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedAt   int64         `json:"created_at"`
	CreatedBy   uuid.UUID     `json:"created_by"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
}

type GetArticlesByUserIDParams struct {
	UserID    uuid.UUID `json:"user_id"`
	Search    string    `json:"search"`
	OrderBy   string    `json:"order_by"`
	SqlOffset int32     `json:"sql_offset"`
	SqlLimit  int32     `json:"sql_limit"`
}

type GetArticlesByUserIDRow struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedAt   int64         `json:"created_at"`
	CreatedBy   uuid.UUID     `json:"created_by"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
}

type UpdateArticleParams struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
	ArticleID   uuid.UUID     `json:"article_id"`
}
