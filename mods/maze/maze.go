package maze

import (
	"Adesubomi/backend-2-challenge-golang/mods/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func getMazes(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "List of all Mazes"},
	)
}

func storeMaze(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Maze has been created."},
	)
}

func getMazeSolution(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Shortest path of Maze"},
	)
}

func Bootstrap(f *fiber.App) {
	f.Post("/maze", user.Protected(), storeMaze)
	f.Get("/maze", user.Protected(), getMazes)
	f.Get("/maze/:id/solution", user.Protected(), getMazeSolution)
}
