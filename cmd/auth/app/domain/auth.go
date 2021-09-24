package domain

import "context"

type AuthenticationUsecase interface {
	Register(context.Context, *RegisterRequest) error
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

type AuthenticationRepository interface {
	GetUserByEmail(ctx context.Context, arg *GetUserByEmailParams) (*UserRow, error)
	GetUserByUserId(ctx context.Context, arg *GetUserByUserIdParams) (*UserRow, error)
	InsertAuthData(ctx context.Context, arg *InsertAuthDataParams) error
}

type SessionRepository interface {
	GetSessionByAccessToken(ctx context.Context, accessToken string) (*GetSessionRow, error)
	InsertSession(ctx context.Context, arg *InsertSessionParams) error
}
