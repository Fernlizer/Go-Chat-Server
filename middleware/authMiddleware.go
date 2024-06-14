// middleware/authMiddleware.go
package middleware

import (
	"gochatserver/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a middleware function to validate JWT token
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the token from the Authorization header
	token := c.Get("Authorization")

	// Validate the token
	valid, err := utils.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{Status: false, Data: "Error validating token"})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{Status: false, Data: "Invalid token"})
	}

	// Move to the next middleware or route handler

	return c.Next()
}
