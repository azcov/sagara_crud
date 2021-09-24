package http

import (
	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
	"github.com/azcov/sagara_crud/cmd/auth/app/interfaces/http/handler"
	appMiddleware "github.com/azcov/sagara_crud/internal/http/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(usecase domain.AuthenticationUsecase) *echo.Echo {
	handler := handler.NewHandler(usecase)
	router := echo.New()

	// Middleware
	router.Use(appMiddleware.EchoCORS())
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.Use(middleware.BodyDump(appMiddleware.DumpRequestResponse))
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	routerAuth := router.Group("/auth")
	routerAuth.POST("/register", handler.Register)
	routerAuth.POST("/login", handler.Login)

	return router
}
