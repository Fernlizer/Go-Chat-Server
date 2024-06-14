// handler/book_handler.go
package handlers

import (
	"gochatserver/utils"

	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
	return c.JSON(utils.ApiResponse{Status: true, Data: nil})
}
