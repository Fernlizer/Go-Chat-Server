// middleware/cors_middleware.go
package middleware

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
)

// ValidateDomainMiddleware is a middleware function to check if the request origin is in the list of allowed domains.
func ValidateDomainMiddleware(c *fiber.Ctx) error {

	// Get the Origin header from the request
	origin := c.IP()
	fmt.Println("Origin:", origin)
	// Get the allowed domains from the configuration
	allowedDomains := viper.GetStringSlice("allowed_domains")

	// Check if the origin is in the allowed domains list
	if !isDomainAllowed(origin, allowedDomains) {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"status": false,
			"data":   "Domain not allowed",
		})
	}

	// Set the necessary headers for CORS
	c.Set("Access-Control-Allow-Origin", origin)
	c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	c.Set("Access-Control-Allow-Credentials", "true")

	// Continue to the next middleware or route handler
	return c.Next()
}

func isDomainAllowed(domain string, allowedDomains []string) bool {
	for _, allowed := range allowedDomains {
		// fmt.Println(domain)
		// fmt.Println(allowed)
		if domain == allowed {
			return true
		}
	}
	return false
}
