package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/azcov/sagara_crud/cmd/product/app/domain"
	appConfig "github.com/azcov/sagara_crud/cmd/product/config"
	"github.com/google/uuid"
)

type productUsecase struct {
	repository domain.ProductRepository
	config     *appConfig.Config
}

func NewUsecase(cfg *appConfig.Config, repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		repository: repo,
		config:     cfg,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, req *domain.CreateProductRequest, userID uuid.UUID) error {
	params := &domain.InsertProductParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Qty:         req.Qty,
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   userID,
	}
	if err := u.repository.InsertProduct(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) GetProduct(ctx context.Context, req *domain.GetProductRequest) (*domain.GetProductResponse, error) {
	data, err := u.repository.GetProductByProductID(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}

	res := &domain.GetProductResponse{
		ProductID:   data.ID.String(),
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Qty:         data.Qty,
		CreatedAt:   data.CreatedAt,
		CreatedBy:   data.CreatedBy.String(),
	}

	if data.UpdatedAt.Valid {
		res.UpdatedAt = &data.UpdatedAt.Int64
	}
	if data.UpdatedBy.Valid {
		updatedBy := data.UpdatedBy.UUID.String()
		res.UpdatedBy = &updatedBy
	}
	return res, nil
}

func (u *productUsecase) GetProducts(ctx context.Context, req *domain.GetProductsRequest, userID uuid.UUID) ([]*domain.GetProductResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	params := &domain.GetProductsByUserIDParams{
		UserID:   userID,
		Search:   req.Search,
		OrderBy:  req.SortBy,
		SqlLimit: req.Limit,
	}
	data, err := u.repository.GetProductsByUserID(ctx, params, req.Page)
	if err != nil {
		return nil, err
	}

	res := []*domain.GetProductResponse{}
	for i := range data {
		product := &domain.GetProductResponse{
			ProductID:   data[i].ID.String(),
			Name:        data[i].Name,
			Description: data[i].Description,
			Price:       data[i].Price,
			Qty:         data[i].Qty,
			CreatedAt:   data[i].CreatedAt,
			CreatedBy:   data[i].CreatedBy.String(),
		}

		if data[i].UpdatedAt.Valid {
			product.UpdatedAt = &data[i].UpdatedAt.Int64
		}
		if data[i].UpdatedBy.Valid {
			updatedBy := data[i].UpdatedBy.UUID.String()
			product.UpdatedBy = &updatedBy
		}
		res = append(res, product)
	}

	return res, nil
}

func (u *productUsecase) UpdateProduct(ctx context.Context, req *domain.UpdateProductRequest, userID uuid.UUID) error {
	params := &domain.UpdateProductParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Qty:         req.Qty,
		UpdatedAt:   sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		UpdatedBy:   uuid.NullUUID{UUID: userID, Valid: true},
		ProductID:   req.ProductID,
	}
	if err := u.repository.UpdateProduct(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) DeleteProduct(ctx context.Context, req *domain.DeleteProductRequest, userID uuid.UUID) error {
	params := &domain.DeleteProductParams{
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		DeletedBy: uuid.NullUUID{UUID: userID, Valid: true},
		ProductID: req.ProductID,
	}
	if err := u.repository.DeleteProduct(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *productUsecase) ForceDeleteProduct(ctx context.Context, req *domain.ForceDeleteProductRequest) error {
	if err := u.repository.ForceDeleteProduct(ctx, req.ProductID); err != nil {
		return err
	}
	return nil
}
