package maze

import (
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
	f.Post("/maze", storeMaze)
	f.Get("/maze", getMazes)
	f.Get("/maze/:id/solution", getMazeSolution)
}
