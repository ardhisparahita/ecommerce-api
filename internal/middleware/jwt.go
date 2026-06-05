package middleware

import (
	"strings"

	"github.com/ardhisparahita/ecommerce-api/pkg/config"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "unauthorized",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := utils.ValidateToken(tokenString, config.Get("JWT_SECRET"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Locals(
			"user_id",
			uint64(claims["user_id"].(float64)),
		)

		return c.Next()
	}
}
