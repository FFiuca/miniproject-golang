package middleware

import (
	"fmt"
	"project1/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(config.Secret),
		},
		ContextKey:  "user",
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		SuccessHandler: func(c *fiber.Ctx) error {
			fmt.Println(c.Locals("user"))

			// normalize user metadata
			userToken := c.Locals("user").(*jwt.Token) // convert to token
			claims := userToken.Claims.(jwt.MapClaims)

			c.Locals("metaUser", claims) // store to request
			fmt.Println("claim", claims)
			return c.Next()
		},
	})
}
