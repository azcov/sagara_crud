package handler

import (
	"github.com/azcov/sagara_crud/cmd/article/app/domain"
	httpUtil "github.com/azcov/sagara_crud/internal/http/echo/util"
	"github.com/azcov/sagara_crud/pkg/validation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	usecase domain.ArticleUsecase
}

func NewArticleHandler(u domain.ArticleUsecase) *articleHandler {
	return &articleHandler{
		usecase: u,
	}
}

// CreateArticle godoc
// @Summary CreateArticle User
// @Description CreateArticle User
// @Tags Article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param request body domain.CreateArticleRequest true "Request Body"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /articles [post]
func (h *articleHandler) CreateArticle(c echo.Context) error {
	req := domain.CreateArticleRequest{}
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

	err = h.usecase.CreateArticle(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.CreatedResponse(c, "success create", nil)
}

// GetArticle godoc
// @Summary GetArticle User
// @Description GetArticle User
// @Tags Article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param article_id path string true "article id"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /articles/{article_id} [get]
func (h *articleHandler) GetArticle(c echo.Context) error {
	req := domain.GetArticleRequest{}
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

	res, err := h.usecase.GetArticle(ctx, &req)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success get article", res)
}

// GetArticles godoc
// @Summary GetArticles User
// @Description GetArticles User
// @Tags Article
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
// @Router /articles [get]
func (h *articleHandler) GetArticles(c echo.Context) error {
	req := domain.GetArticlesRequest{}
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

	res, err := h.usecase.GetArticles(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success get article", res)
}

// UpdateArticle godoc
// @Summary UpdateArticle User
// @Description UpdateArticle User
// @Tags Article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param article_id path string true "article id"
// @Param request body domain.UpdateArticleRequest true "Request Body"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /articles/{article_id} [put]
func (h *articleHandler) UpdateArticle(c echo.Context) error {
	req := domain.UpdateArticleRequest{}
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

	err = h.usecase.UpdateArticle(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success update article", nil)
}

// DeleteArticle godoc
// @Summary DeleteArticle User
// @Description DeleteArticle User
// @Tags Article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param article_id path string true "article id"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /articles/{article_id} [delete]
func (h *articleHandler) DeleteArticle(c echo.Context) error {
	req := domain.DeleteArticleRequest{}
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

	err = h.usecase.DeleteArticle(ctx, &req, userID)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success delete article", nil)
}

// ForceDeleteArticle godoc
// @Summary ForceDeleteArticle User
// @Description ForceDeleteArticle User
// @Tags Article
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer (token)"
// @Param article_id path string true "article id"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /articles/{article_id} [delete]
func (h *articleHandler) ForceDeleteArticle(c echo.Context) error {
	req := domain.ForceDeleteArticleRequest{}
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

	err = h.usecase.ForceDeleteArticle(ctx, &req)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success force delete article", nil)
}
