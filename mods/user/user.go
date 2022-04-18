package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func registerUser(c *fiber.Ctx) error {

	type UserData struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	var ud UserData

	if err := c.BodyParser(&ud); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to create new user"},
		)
	}

	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "User account has been created"},
	)
}

func Bootstrap(f *fiber.App) {
	f.Post("/login", login)
	f.Post("/user", registerUser)
}
