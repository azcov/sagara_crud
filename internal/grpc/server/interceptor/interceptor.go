package interceptor

import (
	"context"
	"strings"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
	commonError "github.com/azcov/sagara_crud/errors"
	"github.com/azcov/sagara_crud/internal/authentication"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var ValidationTokenMethod = map[string]bool{
	"ValidationToken": true,
}
var ValidationRefreshTokenMethod = map[string]bool{
	"ValidationRefreshToken": true,
}

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	repository    domain.SessionRepository
	authenticator authentication.Authenticator
	log           *zap.SugaredLogger
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(a authentication.Authenticator, repo domain.SessionRepository, log *zap.SugaredLogger) *AuthInterceptor {
	return &AuthInterceptor{
		repository:    repo,
		authenticator: a,
		log:           log,
	}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		i.log.Named("interceptor.unary.method").Info(info.FullMethod)

		splittedMethod := strings.Split(info.FullMethod, "/")
		method := splittedMethod[len(splittedMethod)-1]
		if ValidationTokenMethod[method] {
			identity, err := i.validationToken(ctx)
			if err != nil {
				i.log.Named("interceptor.unary.error").Error(err.Error())
				return nil, err
			}
			ctx = context.WithValue(ctx, MetadataKeyUserID, identity.UserID.String())
			ctx = context.WithValue(ctx, MetadataKeyRoleID, identity.RoleID)
		}

		if ValidationRefreshTokenMethod[method] {
			identity, err := i.validationRefreshToken(ctx)
			if err != nil {
				i.log.Named("interceptor.unary.error").Error(err.Error())
				return nil, err
			}
			ctx = context.WithValue(ctx, MetadataKeyUserID, identity.UserID.String())
			ctx = context.WithValue(ctx, MetadataKeyRoleID, identity.RoleID)
		}

		return handler(ctx, req)
	}
}

// --- validationToken used to --- //
func (i *AuthInterceptor) validationToken(ctx context.Context) (*authentication.Identity, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrAuthorizationNotProvided.Error())
	}

	authorization := md[MetadataKeyAuthorization.String()]
	if len(authorization) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrAuthorizationNotProvided.Error())
	}

	token := authorization[0]
	jwtToken, err := i.authenticator.VerifyAccessToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrAuthorizationNotProvided.Error())
	}

	identity, err := i.authenticator.ExtractTokenMetadata(jwtToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrExpiredAuthorization.Error())
	}

	return identity, nil
}

// --- validationRefreshToken used to --- //
func (i *AuthInterceptor) validationRefreshToken(ctx context.Context) (*authentication.Identity, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrAuthorizationNotProvided.Error())
	}

	authorization := md[MetadataKeyAuthorization.String()]
	if len(authorization) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrAuthorizationNotProvided.Error())
	}

	token := authorization[0]

	jwtToken, err := i.authenticator.VerifyRefreshToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrAuthorizationNotProvided.Error())
	}

	identity, err := i.authenticator.ExtractTokenMetadata(jwtToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, commonError.ErrExpiredAuthorization.Error())
	}

	return identity, nil
}
