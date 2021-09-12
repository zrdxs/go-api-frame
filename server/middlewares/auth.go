package middlewares

import (
	"errors"

	"github.com/MarceloZardoBR/go-api-frame/infra/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserTokenKey string

var (
	UserToken UserTokenKey = "UserToken"
)

// TODO: Error Wrap
// Authorization middleware for JWT Token
func Authorization() fiber.Handler {
	return func(fiberCtx *fiber.Ctx) error {

		tokenString := fiberCtx.Get(auth.TokenHeader)

		err := auth.ValidateTokenHeader(tokenString)
		if err != nil {
			return err
		}

		token, ok := auth.GetTokenFromHeader(tokenString).(*jwt.Token)
		if !ok {
			return errors.New("Invalid Token")
		}

		if !token.Valid {
			return errors.New("Invalid Token")
		}

		return fiberCtx.Next()
	}
}