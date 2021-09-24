package errors

import (
	"fmt"
	"net/http"
)

var commonErrorMap = map[error]int{
	ErrEmailNotFound:       http.StatusNotFound,
	ErrPasswordDoesntMatch: http.StatusBadRequest,
	ErrProductNotFound:     http.StatusNotFound,
}

// CommonError is
func CommonError(err error) (int, error) {
	if status, ok := commonErrorMap[err]; ok {
		return status, err
	}

	return http.StatusInternalServerError, fmt.Errorf("internal error")
}
