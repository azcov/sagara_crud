package util

import "github.com/labstack/echo/v4"

func ParsingParameter(ctx echo.Context, i interface{}) error {
	err := ctx.Bind(i)
	if err != nil {
		return &ParsingError{err.Error()}
	}
	return err
}
