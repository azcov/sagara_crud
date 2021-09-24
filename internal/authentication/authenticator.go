package authentication

import (
	"strings"
	"time"

	commonError "github.com/azcov/sagara_crud/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	log "go.uber.org/zap"
)

var (
	jwtSigningMethod = jwt.SigningMethodHS256
)

type JWTToken struct {
	jwt.StandardClaims
	Identity
}

type Identity struct {
	RoleID int32     `json:"role_id"`
	UserID uuid.UUID `json:"user_id"`
}

type Token struct {
	SignedToken string
	ExpiresAt   time.Time
}

type Authenticator interface {
	GenerateAccessToken(userID uuid.UUID, roleID int32) (*Token, error)
	GenerateRefreshToken(userID uuid.UUID, roleID int32) (*Token, error)
	VerifyAccessToken(authorization string) (*jwt.Token, error)
	VerifyRefreshToken(authorization string) (*jwt.Token, error)
	ExtractClaims(tokenStr string) (*Identity, bool)
	ExtractTokenMetadata(token *jwt.Token) (*Identity, error)
}
type secretAuthenticator struct {
	AccessSecret  string
	RefreshSecret string
	tokenType     string
	appName       string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewSecretAuthenticator(tokenType, appName, accessSecret, refreshSecret string, accessExpiry, refreshExpiry time.Duration) Authenticator {
	return secretAuthenticator{
		tokenType:     tokenType,
		appName:       appName,
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (a secretAuthenticator) GenerateAccessToken(userID uuid.UUID, roleID int32) (*Token, error) {
	return a.generate(userID, roleID, a.AccessSecret, a.accessExpiry)
}

func (a secretAuthenticator) GenerateRefreshToken(userID uuid.UUID, roleID int32) (*Token, error) {
	return a.generate(userID, roleID, a.RefreshSecret, a.refreshExpiry)
}

func (a secretAuthenticator) VerifyAccessToken(authorization string) (*jwt.Token, error) {
	return a.verify(authorization, a.AccessSecret)
}

func (a secretAuthenticator) VerifyRefreshToken(authorization string) (*jwt.Token, error) {
	return a.verify(authorization, a.RefreshSecret)
}

// ExtractClaims is
func (a secretAuthenticator) ExtractClaims(tokenStr string) (*Identity, bool) {
	hmacSecretString := viper.GetString("email.token_secret")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uuid.MustParse(claims["sub"].(string))
		roleID := claims["role_id"].(int32)
		return &Identity{
			UserID: userID,
			RoleID: roleID,
		}, true
	}
	log.S().Error("Invalid JWT Token")
	return nil, false
}

func (a secretAuthenticator) ExtractTokenMetadata(token *jwt.Token) (*Identity, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		zap.S().Named("authorization.extract-token.error").Error(commonError.ErrInvalidAuthorization)
		return nil, commonError.ErrInvalidAuthorization
	}
	roleId, ok := claims["role_id"].(float64)
	if !ok {
		zap.S().Named("authorization.extract-token.error").Error(commonError.ErrInvalidAuthorization)
		return nil, commonError.ErrInvalidAuthorization
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		zap.S().Named("authorization.extract-token.error").Error(commonError.ErrInvalidAuthorization)
		return nil, commonError.ErrInvalidAuthorization
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		zap.S().Named("authorization.extract-token.error").Error(commonError.ErrInvalidAuthorization)
		return nil, commonError.ErrInvalidAuthorization
	}
	return &Identity{
		RoleID: int32(roleId),
		UserID: userUUID,
	}, nil

}

func (a secretAuthenticator) generate(userID uuid.UUID, roleID int32, secretKey string, expiry time.Duration) (*Token, error) {
	exp := time.Now().UTC().Add(expiry)
	claims := JWTToken{
		StandardClaims: jwt.StandardClaims{
			Issuer:    a.appName,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Subject:   userID.String(),
		},
		Identity: Identity{
			RoleID: roleID,
			UserID: userID,
		},
	}

	tkn := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err := tkn.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &Token{
		SignedToken: signedToken,
		ExpiresAt:   exp,
	}, err
}

func (a secretAuthenticator) verify(authorization, secretKey string) (*jwt.Token, error) {
	authToken := strings.Split(authorization, " ")
	if len(authToken) != 2 {
		return nil, commonError.ErrInvalidAuthorization
	}

	if authToken[0] != a.tokenType {
		return nil, commonError.ErrInvalidAuthorization
	}
	token := authToken[1]
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, commonError.ErrInvalidAuthorization
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		zap.S().Named("authorization.verify-token.error").Error(err)
		return nil, err
	}

	return jwtToken, nil
}
