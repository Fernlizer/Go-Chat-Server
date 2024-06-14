// routes/routes.go
package routes

import (
	"gochatserver/handlers"
	"gochatserver/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// ChatRoom represents a chat room with a name, a password and a list of users
type ChatRoom struct {
	Name     string
	Password string
	Users    []*websocket.Conn
}

// ChatServer manages the chat rooms and the admin
type ChatServer struct {
	Rooms map[string]*ChatRoom
	Admin *websocket.Conn
}

// declare a global variable for the chat server
var cs *ChatServer

func SetupRoutes(app *fiber.App) {
	// Define user routes
	userRoutes := app.Group("/user")
	userRoutes.Post("/register", handlers.Register)
	userRoutes.Post("/login", handlers.Login)
	userRoutes.Post("/validate", handlers.ValidateToken)

	// Add more routes as needed
	bookRoute := app.Group("/book")
	bookRoute.Use(middleware.ValidateDomainMiddleware)
	bookRoute.Use(middleware.AuthMiddleware)
	bookRoute.Post("/insert", handlers.Insert)

}
