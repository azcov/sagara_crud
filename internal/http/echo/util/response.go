package util

import (
	"fmt"
	"net/http"
	"time"

	errorUtil "github.com/azcov/sagara_crud/errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ResponseStatus string

const (
	ResponseStatusSuccessText             ResponseStatus = "success"
	ResponseStatusCreatedText             ResponseStatus = "success insert data"
	ResponseStatusConflictText            ResponseStatus = "conflict"
	ResponseStatusInternalServerErrorText ResponseStatus = "internal server error"
	ResponseStatusBadRequestText          ResponseStatus = "bad request"
	ResponseStatusNotFoundText            ResponseStatus = "not found"
	ResponseStatusUnprocessableEntityText ResponseStatus = "unprocessable entity"
	ResponseStatusUnauthorized            ResponseStatus = "unauthorized"
	ResponseStatusForbidden               ResponseStatus = "forbidden"
)

//ParsingError is
type ParsingError struct {
	msg string
}

func (re *ParsingError) Error() string { return re.msg }

type Base struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}

// SuccessResponse returns
func SuccessResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := Base{
		Status:     string(ResponseStatusSuccessText),
		StatusCode: http.StatusOK,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Info("success response")

	return ctx.JSON(http.StatusOK, responseData)
}

// CreatedResponse returns
func CreatedResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := Base{
		Status:     string(ResponseStatusCreatedText),
		StatusCode: http.StatusCreated,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Info("success create data")

	return ctx.JSON(http.StatusCreated, responseData)
}

// ErrorResponse returns
func ErrorResponse(ctx echo.Context, err error, data interface{}) error {
	statusCode, err := errorUtil.CheckErrorType(err)
	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, err, data)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, err, data)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, err, data)
	case http.StatusUnauthorized:
		return ErrorUnauthorized(ctx, err, data)
	case http.StatusForbidden:
		return ErrorForbidden(ctx, err, data)
	}
	return ErrorInternalServerResponse(ctx, err, data)
}

// ErrorConflictResponse returns
func ErrorConflictResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := Base{
		Status:     string(ResponseStatusConflictText),
		StatusCode: http.StatusConflict,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("conflict data error : %s ", err.Error())

	return ctx.JSON(http.StatusConflict, responseData)
}

// ErrorInternalServerResponse returns
func ErrorInternalServerResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := Base{
		Status:     string(ResponseStatusInternalServerErrorText),
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("internal server error : %s ", err.Error())

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorBadRequest returns
func ErrorBadRequest(ctx echo.Context, err error, data interface{}) error {
	responseData := Base{
		Status:     string(ResponseStatusBadRequestText),
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("bad request error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorNotFound returns
func ErrorNotFound(ctx echo.Context, err error, data interface{}) error {
	responseData := Base{
		Status:     string(ResponseStatusNotFoundText),
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("error not found : %s ", err.Error())

	return ctx.JSON(http.StatusNotFound, responseData)
}

// ErrorParsing returns
func ErrorParsing(ctx echo.Context, err error, data interface{}) error {

	responseData := Base{
		Status:     string(ResponseStatusUnprocessableEntityText),
		StatusCode: http.StatusUnprocessableEntity,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("parsing data error : %s ", err.Error())

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ErrorValidate returns
func ErrorValidate(ctx echo.Context, err error, data interface{}) error {
	message := errorUtil.ValidationErrors(err)
	responseData := Base{
		Status:     string(ResponseStatusBadRequestText),
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("validate data error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorUnauthorized returns
func ErrorUnauthorized(ctx echo.Context, err error, data interface{}) error {
	responseData := Base{
		Status:     string(ResponseStatusUnauthorized),
		StatusCode: http.StatusUnauthorized,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("unauthorized error : %s ", err.Error())

	return ctx.JSON(http.StatusUnauthorized, responseData)
}

// ErrorForbidden returns
func ErrorForbidden(ctx echo.Context, err error, data interface{}) error {
	responseData := Base{
		Status:     string(ResponseStatusForbidden),
		StatusCode: http.StatusForbidden,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	zap.S().Errorf("forbidden error : %s ", err.Error())

	return ctx.JSON(http.StatusForbidden, responseData)
}

// ErrorDefaultResponse returns
func ErrorDefaultResponse(ctx echo.Context, statusCode int, message string) error {

	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, fmt.Errorf(message), nil)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, fmt.Errorf(message), nil)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, fmt.Errorf(message), nil)
	case http.StatusUnauthorized:
		return ErrorUnauthorized(ctx, fmt.Errorf(message), nil)
	}
	return ErrorInternalServerResponse(ctx, fmt.Errorf(message), nil)
}
