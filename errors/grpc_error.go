package errors

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// * isRpcError used to check error if error is rpc error
func isRpcError(err error) bool {
	if _, ok := status.FromError(err); ok {
		return true
	}
	return false
}

var rpcErrorMap = map[codes.Code]int{
	codes.Canceled:           http.StatusRequestTimeout,
	codes.Unknown:            http.StatusBadRequest,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusRequestTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusRequestTimeout,
	codes.FailedPrecondition: http.StatusBadRequest,
	codes.Aborted:            http.StatusRequestTimeout,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.Internal:           http.StatusInternalServerError,
}

// RpcError is
func RpcError(err error) (int, error) {
	if st, ok := status.FromError(err); ok {
		return rpcErrorMap[st.Code()], fmt.Errorf(st.Message())
	}

	return http.StatusInternalServerError, fmt.Errorf("internal error")
}

func RpcErrorCodeIs(err error, code codes.Code) bool {
	if st, ok := status.FromError(err); ok {
		return st.Code() == code
	}

	return false
}
