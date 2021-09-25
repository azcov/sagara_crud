package domain

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
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

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req *CreateProductRequest, userID uuid.UUID) error
	GetProduct(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error)
	GetProducts(ctx context.Context, req *GetProductsRequest, userID uuid.UUID) ([]*GetProductResponse, error)
	UpdateProduct(ctx context.Context, req *UpdateProductRequest, userID uuid.UUID) error
	DeleteProduct(ctx context.Context, req *DeleteProductRequest, userID uuid.UUID) error
	ForceDeleteProduct(ctx context.Context, req *ForceDeleteProductRequest) error
}

type ProductRepository interface {
	InsertProduct(ctx context.Context, arg *InsertProductParams) error
	GetProductsByUserID(ctx context.Context, arg *GetProductsByUserIDParams, page int32) ([]*Product, error)
	GetProductByProductID(ctx context.Context, productID uuid.UUID) (*Product, error)
	UpdateProduct(ctx context.Context, arg *UpdateProductParams) error
	DeleteProduct(ctx context.Context, arg *DeleteProductParams) error
	ForceDeleteProduct(ctx context.Context, productID uuid.UUID) error
}
