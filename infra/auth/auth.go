package auth

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var privateKey string

type AuthToken struct {
	bar string `json:"bar"`
	jwt.StandardClaims
}

// GenerateToken for app
func (s *AuthToken) GenerateToken() (appToken string, err error) {

	var key *rsa.PrivateKey
	key, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return appToken, err
	}

	claims := AuthToken{}
	claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
	claims.Issuer = "entity-z"

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	appToken, err = token.SignedString(key)
	if err != nil {
		return appToken, err
	}

	return appToken, nil
}
