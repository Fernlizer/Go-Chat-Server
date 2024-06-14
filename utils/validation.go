// utils/validation.go
package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateAndParseBody validates the request body and parses it into the given structure
func ValidateAndParseBody(c *fiber.Ctx, v interface{}) error {
	if err := c.BodyParser(v); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
	}

	if err := validate.Struct(v); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

var jwtSecret = []byte("FERNSERVSECRETEKEY")

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// ValidateToken validates a JWT token and returns the validation status.
func ValidateToken(tokenString string) (bool, error) {
	claims := &Claims{}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}

func DecodeTokenClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	if _, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("Decoded Claims: %+v\n", claims)
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}
