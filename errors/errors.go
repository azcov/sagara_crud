package errors

import "fmt"

var (
	ErrAuthorizationNotProvided = fmt.Errorf("authorization token is not provided")
	ErrInvalidAuthorization     = fmt.Errorf("invalid authorization")
	ErrExpiredAuthorization     = fmt.Errorf("authorization token is expired")

	ErrEmailNotFound       = fmt.Errorf("email not found")
	ErrPasswordDoesntMatch = fmt.Errorf("password doesnt match")
	ErrArticleNotFound     = fmt.Errorf("article not found")
)

func CheckErrorType(err error) (int, error) {
	switch {
	case isPqError(err):
		return PqError(err)
	case isRpcError(err):
		return RpcError(err)
	}
	return CommonError(err)
}
