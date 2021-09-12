package auth

import (
	"crypto/rsa"
	"net/http"
	"strings"
	"time"

	"github.com/MarceloZardoBR/go-api-frame/infra/config"
	"github.com/MarceloZardoBR/go-api-frame/server/viewmodels"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

const (
	TokenHeader string = "Token"
)

type AuthToken struct {
	bar string `json:"bar"`
	jwt.StandardClaims
}

func (s *AuthToken) GenerateUserTokenAndResponse(cfg *config.Config) (authResponse viewmodels.AuthResponse, err error) {

	IssuedAt := time.Now().Unix()
	ExpiresAt := time.Now().Add(time.Hour * time.Duration(cfg.AuthTokenExpireTime)).Unix()

	claims := AuthToken{}
	claims.IssuedAt = IssuedAt
	claims.ExpiresAt = ExpiresAt
	claims.Issuer = "entity-z"

	token, err := s.GenerateToken(cfg, claims)
	if err != nil {
		return authResponse, errors.WithStack(err)
	}

	authResponse.Token = token
	authResponse.ExpirationTime = ExpiresAt

	return authResponse, nil
}

// GenerateToken for app
func (s *AuthToken) GenerateToken(cfg *config.Config, claims jwt.Claims) (appToken string, err error) {
	var key *rsa.PrivateKey
	key, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(cfg.APIPrivateKey))
	if err != nil {
		return appToken, errors.WithStack(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	appToken, err = token.SignedString(key)
	if err != nil {
		return appToken, errors.WithStack(err)
	}

	return appToken, nil
}

// GetTokenData returns token data
func (s *AuthToken) GetTokenData(tokenString string, cfg config.Config) (tokenData AuthToken, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(http.StatusText(http.StatusForbidden))
		}

		return cfg.APIPrivateKey, nil
	})

	if tokenData, ok := token.Claims.(AuthToken); !ok {
		return tokenData, errors.New(http.StatusText(http.StatusForbidden))
	}

	return tokenData, nil
}

// GetTokenFromHeader returns token from header
func GetTokenFromHeader(value string) interface{} {
	parts := strings.Split(value, " ")

	return parts[0]
}

// ValidateTokenHeader validate token string
func ValidateTokenHeader(value string) (err error) {
	if strings.TrimSpace(value) == "" {
		return errors.New("Token não encontrado")
	}

	parts := strings.Split(value, " ")

	if len(parts) != 2 || strings.Compare(parts[0], "Bearer") != 0 {
		return errors.New("Token inválido")
	}

	return nil
}
