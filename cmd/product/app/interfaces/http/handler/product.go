package handler

import (
	"github.com/azcov/sagara_crud/cmd/product/app/domain"
	httpUtil "github.com/azcov/sagara_crud/internal/http/echo/util"
	"github.com/azcov/sagara_crud/pkg/validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(u domain.ProductUsecase) *productHandler {
	return &productHandler{
		usecase: u,
	}
}

// CreateProduct godoc
// @Summary CreateProduct User
// @Description CreateProduct User
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param request body domain.CreateProductRequest true "Request Body"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /products [post]
func (h *productHandler) CreateProduct(c echo.Context) error {
	req := domain.CreateProductRequest{}
	ctx := c.Request().Context()
	userID := c.Get("user_id").(uuid.UUID)

	// parsing
	err := httpUtil.ParsingParameter(c, &req)
	if err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}

	// validate
	err = validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	err = h.usecase.CreateProduct(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.CreatedResponse(c, "success create", nil)
}

// GetProduct godoc
// @Summary GetProduct User
// @Description GetProduct User
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param product_id path string true "product id"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /products/{product_id} [get]
func (h *productHandler) GetProduct(c echo.Context) error {
	req := domain.GetProductRequest{}
	ctx := c.Request().Context()

	// parsing
	err := httpUtil.ParsingParameter(c, &req)
	if err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}

	// validate
	err = validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	res, err := h.usecase.GetProduct(ctx, &req)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success get product", res)
}

// GetProducts godoc
// @Summary GetProducts User
// @Description GetProducts User
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param sort_by query string false "sort by"
// @Param search query string false "search"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /products [get]
func (h *productHandler) GetProducts(c echo.Context) error {
	req := domain.GetProductsRequest{}
	ctx := c.Request().Context()
	userID := c.Get("user_id").(uuid.UUID)

	// parsing
	err := httpUtil.ParsingParameter(c, &req)
	if err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}

	// validate
	err = validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	res, err := h.usecase.GetProducts(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success get product", res)
}

// UpdateProduct godoc
// @Summary UpdateProduct User
// @Description UpdateProduct User
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param product_id path string true "product id"
// @Param request body domain.UpdateProductRequest true "Request Body"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /products/{product_id} [put]
func (h *productHandler) UpdateProduct(c echo.Context) error {
	req := domain.UpdateProductRequest{}
	ctx := c.Request().Context()
	userID := c.Get("user_id").(uuid.UUID)

	// parsing
	err := httpUtil.ParsingParameter(c, &req)
	if err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}

	// validate
	err = validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	err = h.usecase.UpdateProduct(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success update product", nil)
}

// DeleteProduct godoc
// @Summary DeleteProduct User
// @Description DeleteProduct User
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param product_id path string true "product id"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /products/{product_id} [delete]
func (h *productHandler) DeleteProduct(c echo.Context) error {
	req := domain.DeleteProductRequest{}
	ctx := c.Request().Context()
	userID := c.Get("user_id").(uuid.UUID)

	// parsing
	err := httpUtil.ParsingParameter(c, &req)
	if err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}

	// validate
	err = validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	err = h.usecase.DeleteProduct(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success delete product", nil)
}

// ForceDeleteProduct godoc
// @Summary ForceDeleteProduct User
// @Description ForceDeleteProduct User
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param product_id path string true "product id"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /products/{product_id} [delete]
func (h *productHandler) ForceDeleteProduct(c echo.Context) error {
	req := domain.ForceDeleteProductRequest{}
	ctx := c.Request().Context()

	// parsing
	err := httpUtil.ParsingParameter(c, &req)
	if err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}

	// validate
	err = validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	err = h.usecase.ForceDeleteProduct(ctx, &req)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success force delete product", nil)
}
