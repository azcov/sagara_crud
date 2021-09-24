package errors

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/lib/pq"
)

// * isPqError used to check error if error is pg error
func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	re := regexp.MustCompile("\\((.*?)\\)")
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)
		// Change Field Name
		if len(match) >= 2 {
			switch match[1] {
			}
			switch err.Code.Name() {
			case "unique_violation":
				return pqErrorMap["unique_violation"], fmt.Errorf("%s already exists", match[1])
			}
		}

	}

	return http.StatusInternalServerError, fmt.Errorf("internal error")
}
