// handler/user_handler.go
package handlers

import (
	"gochatserver/db"
	"gochatserver/db/models"
	"gochatserver/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	// Parse request body
	var user models.User
	if err := utils.ValidateAndParseBody(c, &user); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(utils.ApiResponse{Status: false, Data: err.Error()})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{Status: false, Data: "Error hashing the password"})
	}

	// Set the hashed password in the user model
	user.Password = string(hashedPassword)

	// Create user in the database
	if err := db.DB.Create(&user).Error; err != nil {
		// Check if the error contains the substring "Error 1062" (MySQL duplicate entry error)
		if strings.Contains(err.Error(), "Error 1062") {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ApiResponse{Status: false, Data: "Username or email already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{Status: false, Data: "Error creating user"})
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{Status: false, Data: "Error generating token"})
	}

	// Clear the password field before returning the user data
	user.Password = ""

	// Include the token in the response
	return c.JSON(utils.ApiResponse{Status: true, Data: map[string]interface{}{"user": user, "token": token}})
}

var jwtSecret = []byte("FERNSERVSECRETEKEY")

func generateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time: 1 day

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(c *fiber.Ctx) error {

	var loginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := utils.ValidateAndParseBody(c, &loginRequest); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(utils.ApiResponse{Status: false, Data: err.Error()})
	}

	var user models.User
	if err := db.DB.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{Status: false, Data: "Invalid username or password"})
	}

	// Check if the user is active
	if !user.IsActive {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{Status: false, Data: "User is not active"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.ApiResponse{Status: false, Data: "Invalid username or password"})
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{Status: false, Data: "Error generating token"})
	}

	// Clear the password field before returning the user data
	user.Password = ""

	// Include the token in the response
	return c.JSON(utils.ApiResponse{Status: true, Data: map[string]interface{}{"user": user, "token": token}})
}

func ValidateToken(c *fiber.Ctx) error {
	var validateTokenRequest struct {
		Token string `json:"token" validate:"required"`
	}

	if err := utils.ValidateAndParseBody(c, &validateTokenRequest); err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(utils.ApiResponse{Status: false, Data: err.Error()})
	}

	// Decode the token and get the claims
	claims, err := utils.DecodeTokenClaims(validateTokenRequest.Token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponse{Status: false, Data: "Error decoding token"})
	}

	return c.JSON(utils.ApiResponse{Status: true, Data: map[string]interface{}{"claims": claims}})
}
