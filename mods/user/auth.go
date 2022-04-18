package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {

	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Login successful"},
	)
}
