// main.go
package main

import (
	"fmt"
	"gochatserver/db"
	"gochatserver/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Set default values
	viper.SetDefault("allowed_domains", []string{})
}

func main() {
	// Initialize database
	db.InitDB()
	loadConfig()
	// Create Fiber app
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Authentication Server",
	})

	// Set up routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3441")
}
