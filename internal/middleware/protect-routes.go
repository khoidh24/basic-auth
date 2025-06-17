package middleware

import (
	configs "leanGo/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ProtectRoutes() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized - Missing token")
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return configs.JWTSecret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized - Invalid token")
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("username", claims["name"])
		return c.Next()
	}
}
