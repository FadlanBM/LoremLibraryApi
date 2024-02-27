package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var secret = []byte("AllYourBase")

// APIKeyAuthMiddlewareAdmin Authorization godoc
// @Param Authorization header string true "Bearer {token}" "Authorization header for JWT"

func APIKeyAuthMiddleware(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")

	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or empty Authorization header",
		})
	}

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format. Use 'Bearer' schema",
		})
	}

	return c.Next()
}

func APIKeyAuthMiddlewareMePeminjam(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")

	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or empty Authorization header",
		})
	}

	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}
		c.Locals("borrowers", token.Claims)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format. Use 'Bearer' schema",
		})
	}

	return c.Next()
}
