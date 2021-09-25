package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type RequestDump struct {
	Uri  string      `json:"uri"`
	Body interface{} `json:"body"`
}

type ResponseDump struct {
	StatusCode int         `json:"status"`
	Body       interface{} `json:"body"`
}

//EchoCORS is
func EchoCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderXCSRFToken, echo.HeaderAuthorization, echo.HeaderAccessControlAllowOrigin},
	})
}

func DumpRequestResponse(c echo.Context, reqBody, resBody []byte) {
	var requestData, responseData interface{}

	json.Unmarshal(reqBody, &requestData)
	json.Unmarshal(resBody, &responseData)

	request := RequestDump{
		Uri:  c.Request().URL.Path,
		Body: requestData,
	}
	response := ResponseDump{
		StatusCode: c.Response().Status,
		Body:       responseData,
	}

	jsonRequest, _ := json.Marshal(request)
	jsonResponse, _ := json.Marshal(response)

	zap.S().Named("request").Info(string(jsonRequest))
	zap.S().Named("response").Info(string(jsonResponse))
}
