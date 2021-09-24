package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/azcov/sagara_crud/cmd/article/app/domain"
	commonError "github.com/azcov/sagara_crud/errors"
	helperSQL "github.com/azcov/sagara_crud/pkg/sql"
	"github.com/google/uuid"
)

const insertArticle = `-- name: InsertArticle :exec
INSERT INTO articles(title, description, created_at, created_by)
VALUES ($1, $2, $3, $4)
`

func (q *Queries) InsertArticle(ctx context.Context, arg *domain.InsertArticleParams) error {
	_, err := q.exec(ctx, q.insertArticleStmt, insertArticle,
		arg.Title,
		arg.Description,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	return err
}

const deleteArticle = `-- name: DeleteArticle :exec
UPDATE articles
SET
    deleted_at = $1,
    deleted_by = $2
WHERE 
    id = $3
`

func (q *Queries) DeleteArticle(ctx context.Context, arg *domain.DeleteArticleParams) error {
	_, err := q.exec(ctx, q.deleteArticleStmt, deleteArticle, arg.DeletedAt, arg.DeletedBy, arg.ArticleID)
	return err
}

const forceDeleteArticle = `-- name: ForceDeleteArticle :exec
DELETE FROM articles WHERE id = $1
`

func (q *Queries) ForceDeleteArticle(ctx context.Context, articleID uuid.UUID) error {
	_, err := q.exec(ctx, q.forceDeleteArticleStmt, forceDeleteArticle, articleID)
	return err
}

const getArticleByArticleID = `-- name: GetArticleByArticleID :one
SELECT 
    id,
    title,
    description,
    created_at,
    created_by,
    updated_at,
    updated_by
FROM articles
WHERE id = $1 
	AND deleted_at IS NULL
`

func (q *Queries) GetArticleByArticleID(ctx context.Context, articleID uuid.UUID) (*domain.Article, error) {
	row := q.queryRow(ctx, q.getArticleByArticleIDStmt, getArticleByArticleID, articleID)
	var i domain.Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	if err == sql.ErrNoRows {
		err = commonError.ErrArticleNotFound
		return nil, err
	}
	return &i, err
}

const getArticlesByUserID = `-- name: GetArticlesByUserID :many
SELECT 
    id,
    title,
    description,
    created_at,
    created_by,
    updated_at,
    updated_by
FROM articles
WHERE
    deleted_at IS NULL
    AND created_by = $1
    %s
%s
OFFSET $2
LIMIT $3
`

var getArtilcesSort = map[string]string{
	"created": "created_at",
	"updated": "updated_at",
	"title":   "title",
}

func (q *Queries) GetArticlesByUserID(ctx context.Context, arg *domain.GetArticlesByUserIDParams, page int32) ([]*domain.Article, error) {
	arg.SqlOffset = helperSQL.GenerateOffsetCluse(page, arg.SqlLimit)
	if arg.Search != "" {
		arg.Search = "AND LOWER(title) ILIKE '%" + arg.Search + "%'"
	}
	finalQuery := fmt.Sprintf(getArticlesByUserID, arg.Search, helperSQL.GenerateOrderByClause(arg.OrderBy, getArtilcesSort, "created_by desc"))
	rows, err := q.query(ctx, q.getArticlesByUserIDStmt, finalQuery,
		arg.UserID,
		arg.SqlOffset,
		arg.SqlLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*domain.Article{}
	for rows.Next() {
		var i domain.Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateArticle = `-- name: UpdateArticle :exec
UPDATE articles
SET
    title = $1,
    description = $2,
    updated_at = $3,
    updated_by = $4
WHERE 
    id = $5
`

func (q *Queries) UpdateArticle(ctx context.Context, arg *domain.UpdateArticleParams) error {
	_, err := q.exec(ctx, q.updateArticleStmt, updateArticle,
		arg.Title,
		arg.Description,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ArticleID,
	)
	return err
}
