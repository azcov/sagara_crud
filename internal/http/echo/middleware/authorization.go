package middleware

import (
	authProto "github.com/azcov/sagara_crud/cmd/auth/proto"
	"github.com/azcov/sagara_crud/internal/authentication"
	httpUtil "github.com/azcov/sagara_crud/internal/http/echo/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Middleware struct {
	authenticator authentication.Authenticator
	log           *zap.SugaredLogger
	authClient    authProto.AuthServiceClient
}

// NewMiddleware for
func NewMiddleware(authenticator authentication.Authenticator, log *zap.SugaredLogger, authClient authProto.AuthServiceClient) *Middleware {
	return &Middleware{
		authClient:    authClient,
		authenticator: authenticator,
		log:           log,
	}
}

func (m *Middleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := metadata.AppendToOutgoingContext(c.Request().Context(),
			"Authorization", c.Request().Header.Get("Authorization"),
		)
		result, err := m.authClient.ValidationToken(ctx, &emptypb.Empty{})
		if err != nil {
			return httpUtil.ErrorResponse(c, err, nil)
		}

		userID, err := uuid.Parse(result.UserId)
		if err != nil {
			return httpUtil.ErrorResponse(c, err, nil)
		}

		c.Set("user_id", userID)
		return next(c)
	}
}

// JWTRefreshAuth is
func (m *Middleware) JWTRefreshAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		jwtToken, err := m.authenticator.VerifyRefreshToken(authorization)
		if err != nil {
			return httpUtil.ErrorResponse(c, err, nil)
		}
		identity, ok := m.authenticator.ExtractClaims(jwtToken.Raw)
		if !ok {
			return httpUtil.ErrorResponse(c, err, nil)
		}
		c.Set("user_id", identity.UserID)
		c.Set("role_id", identity.RoleID)
		return next(c)
	}
}

// JWTAccessAuth is
func (m *Middleware) JWTAccessAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")
		jwtToken, err := m.authenticator.VerifyAccessToken(authorization)
		if err != nil {
			return httpUtil.ErrorResponse(c, err, nil)
		}
		identity, ok := m.authenticator.ExtractClaims(jwtToken.Raw)
		if !ok {
			return httpUtil.ErrorResponse(c, err, nil)
		}
		c.Set("user_id", identity.UserID)
		c.Set("role_id", identity.RoleID)
		return next(c)
	}
}
