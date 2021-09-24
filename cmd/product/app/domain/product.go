package domain

import (
	"database/sql"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Qty         int64  `json:"qty"`
}

type GetProductRequest struct {
	ProductID uuid.UUID `param:"product_id" json:"-"`
}
type GetProductsRequest struct {
	Search string `query:"search"`
	Page   int32  `query:"page"`
	Limit  int32  `query:"limit"`
	SortBy string `query:"sort_by"`
}

type UpdateProductRequest struct {
	ProductID   uuid.UUID `param:"product_id" json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Qty         int64     `json:"qty"`
}

type DeleteProductRequest struct {
	ProductID uuid.UUID `param:"product_id" json:"-"`
}

type ForceDeleteProductRequest struct {
	ProductID uuid.UUID `param:"product_id" json:"-"`
}

type GetProductResponse struct {
	ProductID   string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int64   `json:"price"`
	Qty         int64   `json:"qty"`
	CreatedAt   int64   `json:"created_at"`
	CreatedBy   string  `json:"created_by"`
	UpdatedAt   *int64  `json:"updated_at"`
	UpdatedBy   *string `json:"updated_by"`
}

type InsertProductParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Qty         int64     `json:"qty"`
	CreatedAt   int64     `json:"created_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
}

type DeleteProductParams struct {
	DeletedAt sql.NullInt64 `json:"deleted_at"`
	DeletedBy uuid.NullUUID `json:"deleted_by"`
	ProductID uuid.UUID     `json:"product_id"`
}

type GetProductByProductIDRow struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       int64         `json:"price"`
	Qty         int64         `json:"qty"`
	CreatedAt   int64         `json:"created_at"`
	CreatedBy   uuid.UUID     `json:"created_by"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
}

type GetProductsByUserIDParams struct {
	UserID    uuid.UUID `json:"user_id"`
	Search    string    `json:"search"`
	OrderBy   string    `json:"order_by"`
	SqlOffset int32     `json:"sql_offset"`
	SqlLimit  int32     `json:"sql_limit"`
}

type GetProductsByUserIDRow struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       int64         `json:"price"`
	Qty         int64         `json:"qty"`
	CreatedAt   int64         `json:"created_at"`
	CreatedBy   uuid.UUID     `json:"created_by"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
}

type UpdateProductParams struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Price       int64         `json:"price"`
	Qty         int64         `json:"qty"`
	UpdatedAt   sql.NullInt64 `json:"updated_at"`
	UpdatedBy   uuid.NullUUID `json:"updated_by"`
	ProductID   uuid.UUID     `json:"product_id"`
}
