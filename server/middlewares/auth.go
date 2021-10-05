package middlewares

import (
	"github.com/MarceloZardoBR/go-api-frame/infra/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type UserTokenKey string

var (
	UserToken UserTokenKey = "UserToken"
)

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

		// TODO: improve this error
		_, ok = token.Claims.(auth.AuthToken)
		if !ok {
			return errors.New("Unauthorized")
		}

		return fiberCtx.Next()
	}
}
