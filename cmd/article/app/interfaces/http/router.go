package http

import (
	"github.com/azcov/sagara_crud/cmd/article/app/domain"
	"github.com/azcov/sagara_crud/cmd/article/app/interfaces/http/handler"
	"github.com/azcov/sagara_crud/internal/authentication"
	appMiddleware "github.com/azcov/sagara_crud/internal/http/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(authenticator authentication.Authenticator, usecase domain.ArticleUsecase, customMiddleware *appMiddleware.Middleware) *echo.Echo {
	handler := handler.NewArticleHandler(usecase)
	router := echo.New()

	// Middleware
	router.Use(appMiddleware.EchoCORS())
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.Use(middleware.BodyDump(appMiddleware.DumpRequestResponse))

	router.GET("/swagger/*", echoSwagger.WrapHandler)

	routerArticle := router.Group("/articles")
	routerArticle.POST("", handler.CreateArticle, customMiddleware.ValidateToken)
	routerArticle.GET("", handler.GetArticles, customMiddleware.ValidateToken)
	routerArticle.GET("/:article_id", handler.GetArticle, customMiddleware.ValidateToken)
	routerArticle.PUT("/:article_id", handler.UpdateArticle, customMiddleware.ValidateToken)
	routerArticle.DELETE("/:article_id", handler.DeleteArticle, customMiddleware.ValidateToken)
	routerArticle.DELETE("/:article_id/force", handler.ForceDeleteArticle, customMiddleware.ValidateToken)

	return router
}
