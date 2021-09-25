package domain

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedAt   int64         `json:"created_at"`
	CreatedBy   uuid.UUID     `json:"created_by"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
	DeletedAt   sql.NullInt64 `json:"deleted_at"`
	DeletedBy   uuid.NullUUID `json:"deleted_by"`
}

type ArticleUsecase interface {
	CreateArticle(ctx context.Context, req *CreateArticleRequest, userID uuid.UUID) error
	GetArticle(ctx context.Context, req *GetArticleRequest) (*GetArticleResponse, error)
	GetArticles(ctx context.Context, req *GetArticlesRequest, userID uuid.UUID) ([]*GetArticleResponse, error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, userID uuid.UUID) error
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, userID uuid.UUID) error
	ForceDeleteArticle(ctx context.Context, req *ForceDeleteArticleRequest) error
}

type ArticleRepository interface {
	InsertArticle(ctx context.Context, arg *InsertArticleParams) error
	GetArticlesByUserID(ctx context.Context, arg *GetArticlesByUserIDParams, page int32) ([]*Article, error)
	GetArticleByArticleID(ctx context.Context, articleID uuid.UUID) (*Article, error)
	UpdateArticle(ctx context.Context, arg *UpdateArticleParams) error
	DeleteArticle(ctx context.Context, arg *DeleteArticleParams) error
	ForceDeleteArticle(ctx context.Context, articleID uuid.UUID) error
}
