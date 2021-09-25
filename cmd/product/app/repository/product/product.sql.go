package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/azcov/sagara_crud/cmd/product/app/domain"
	commonError "github.com/azcov/sagara_crud/errors"
	helperSQL "github.com/azcov/sagara_crud/pkg/sql"
	"github.com/google/uuid"
)

const insertProduct = `-- name: InsertProduct :exec
INSERT INTO products(name, description, price, qty, created_at, created_by)
VALUES ($1, $2, $3, $4, $5, $6)
`

func (q *Queries) InsertProduct(ctx context.Context, arg *domain.InsertProductParams) error {
	_, err := q.exec(ctx, q.insertProductStmt, insertProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Qty,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
UPDATE products
SET
    deleted_at = $1,
    deleted_by = $2
WHERE 
    id = $3
`

func (q *Queries) DeleteProduct(ctx context.Context, arg *domain.DeleteProductParams) error {
	_, err := q.exec(ctx, q.deleteProductStmt, deleteProduct, arg.DeletedAt, arg.DeletedBy, arg.ProductID)
	return err
}

const forceDeleteProduct = `-- name: ForceDeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) ForceDeleteProduct(ctx context.Context, productID uuid.UUID) error {
	_, err := q.exec(ctx, q.forceDeleteProductStmt, forceDeleteProduct, productID)
	return err
}

const getProductByProductID = `-- name: GetProductByProductID :one
SELECT 
    id,
    name,
    description,
    price,
    qty,
    created_at,
    created_by,
    updated_at,
    updated_by
FROM products
WHERE id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetProductByProductID(ctx context.Context, productID uuid.UUID) (*domain.Product, error) {
	row := q.queryRow(ctx, q.getProductByProductIDStmt, getProductByProductID, productID)
	var i domain.Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Qty,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
	)
	if err == sql.ErrNoRows {
		err = commonError.ErrProductNotFound
		return nil, err
	}
	return &i, err
}

const getProductsByUserID = `-- name: GetProductsByUserID :many
SELECT 
	id,
	name,
	description,
	price,
	qty,
	created_at,
	created_by,
	updated_at,
	updated_by
FROM products
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
	"name":    "name",
}

func (q *Queries) GetProductsByUserID(ctx context.Context, arg *domain.GetProductsByUserIDParams, page int32) ([]*domain.Product, error) {
	arg.SqlOffset = helperSQL.GenerateOffsetCluse(page, arg.SqlLimit)
	if arg.Search != "" {
		arg.Search = "AND LOWER(name) ILIKE '%" + arg.Search + "%'"
	}
	finalQuery := fmt.Sprintf(getProductsByUserID, arg.Search, helperSQL.GenerateOrderByClause(arg.OrderBy, getArtilcesSort, "created_by desc"))
	rows, err := q.query(ctx, q.getProductsByUserIDStmt, finalQuery,
		arg.UserID,
		arg.SqlOffset,
		arg.SqlLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*domain.Product{}
	for rows.Next() {
		var i domain.Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Qty,
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

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET
	name = $1,
	description = $2,
	price = $3,
	qty = $4,
	updated_at = $5,
	updated_by = $6
WHERE id = $7
`

func (q *Queries) UpdateProduct(ctx context.Context, arg *domain.UpdateProductParams) error {
	_, err := q.exec(ctx, q.updateProductStmt, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Qty,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ProductID,
	)
	return err
}
