package handler

import (
	"encoding/json"
	"io/ioutil"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
	httpUtil "github.com/azcov/sagara_crud/internal/http/echo/util"
	"github.com/azcov/sagara_crud/pkg/validation"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type handler struct {
	usecase domain.AuthenticationUsecase
}

func NewHandler(u domain.AuthenticationUsecase) *handler {
	return &handler{
		usecase: u,
	}
}

// Register godoc
// @Summary Register User
// @Description Register User
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "Request Body"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /auth/register [post]
func (h *handler) Register(c echo.Context) error {
	req := domain.RegisterRequest{}
	ctx := c.Request().Context()

	body, _ := ioutil.ReadAll(c.Request().Body)
	zap.S().Info(string(body))

	if err := json.Unmarshal(body, &req); err != nil {
		return httpUtil.ErrorParsing(c, err, nil)
	}
	// parsing
	// err := httpUtil.ParsingParameter(c, &req)
	// if err != nil {
	// 	return httpUtil.ErrorParsing(c, err, nil)
	// }

	// validate
	err := validation.ValidateStruct(&req)
	if err != nil {
		return httpUtil.ErrorValidate(c, err, nil)
	}

	err = h.usecase.Register(ctx, &req)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success register user", nil)
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body domain.LoginRequest true "Request Body"
// @Success 200 {object} httpUtil.Base
// @Failure 400 {object} httpUtil.Base
// @Failure 401 {object} httpUtil.Base
// @Failure 404 {object} httpUtil.Base
// @Failure 500 {object} httpUtil.Base
// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) error {
	req := domain.LoginRequest{}
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

	res, err := h.usecase.Login(ctx, &req)
	if err != nil {
		return httpUtil.ErrorResponse(c, err, nil)
	}

	return httpUtil.SuccessResponse(c, "success login", res)
}
