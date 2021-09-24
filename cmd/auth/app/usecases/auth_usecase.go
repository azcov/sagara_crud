package usecase

import (
	"context"
	"time"

	"github.com/azcov/sagara_crud/cmd/auth/app/domain"
	appConfig "github.com/azcov/sagara_crud/cmd/auth/config"
	commonError "github.com/azcov/sagara_crud/errors"
	"github.com/azcov/sagara_crud/internal/authentication"
	"github.com/azcov/sagara_crud/pkg/util"
)

type authUsecase struct {
	repository    domain.AuthenticationRepository
	config        *appConfig.Config
	authenticator authentication.Authenticator
}

func NewUsecase(cfg *appConfig.Config, repo domain.AuthenticationRepository, a authentication.Authenticator) domain.AuthenticationUsecase {
	return &authUsecase{
		repository:    repo,
		config:        cfg,
		authenticator: a,
	}
}

func (u *authUsecase) Register(ctx context.Context, req *domain.RegisterRequest) error {

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}
	params := domain.InsertAuthDataParams{
		RoleID:      int32(domain.RoleID(req.RoleID)),
		Name:        req.Name,
		Email:       req.Email,
		EmailSecret: u.config.Encryption.EmailPassphrase,
		Password:    hashedPassword,
		CreatedAt:   time.Now().Unix(),
	}
	err = u.repository.InsertAuthData(ctx, &params)
	if err != nil {
		return err
	}
	return nil
}

func (u *authUsecase) Login(ctx context.Context, req *domain.LoginRequest) (*domain.LoginResponse, error) {
	params := &domain.GetUserByEmailParams{
		Email:       []byte(req.Email),
		EmailSecret: u.config.Encryption.EmailPassphrase,
	}

	user, err := u.repository.GetUserByEmail(ctx, params)
	if err != nil {
		return nil, err
	}

	if err = util.CheckPasswordHash(user.Password, req.Password); err != nil {
		err = commonError.ErrPasswordDoesntMatch
		return nil, err
	}

	accessToken, err := u.authenticator.GenerateAccessToken(user.UserID, user.RoleID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.authenticator.GenerateRefreshToken(user.UserID, user.RoleID)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		TokenType:    u.config.Auth.TokenType,
		Token:        accessToken.SignedToken,
		RefreshToken: refreshToken.SignedToken,
		Expires:      int32(u.config.Auth.AccessTokenExpiry.Seconds()),
	}, nil
}
