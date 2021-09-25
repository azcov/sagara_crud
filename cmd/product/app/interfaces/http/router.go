package http

import (
	"github.com/azcov/sagara_crud/cmd/product/app/domain"
	"github.com/azcov/sagara_crud/cmd/product/app/interfaces/http/handler"
	"github.com/azcov/sagara_crud/internal/authentication"
	appMiddleware "github.com/azcov/sagara_crud/internal/http/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(authenticator authentication.Authenticator, usecase domain.ProductUsecase, customMiddleware *appMiddleware.Middleware) *echo.Echo {
	handler := handler.NewProductHandler(usecase)
	router := echo.New()

	// Middleware
	router.Use(appMiddleware.EchoCORS())
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.Use(middleware.BodyDump(appMiddleware.DumpRequestResponse))

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	routerProduct := router.Group("/products")
	routerProduct.POST("", handler.CreateProduct, customMiddleware.ValidateToken)
	routerProduct.GET("", handler.GetProducts, customMiddleware.ValidateToken)
	routerProduct.GET("/:product_id", handler.GetProduct, customMiddleware.ValidateToken)
	routerProduct.PUT("/:product_id", handler.UpdateProduct, customMiddleware.ValidateToken)
	routerProduct.DELETE("/:product_id", handler.DeleteProduct, customMiddleware.ValidateToken)
	routerProduct.DELETE("/:product_id/force", handler.ForceDeleteProduct, customMiddleware.ValidateToken)

	return router
}
