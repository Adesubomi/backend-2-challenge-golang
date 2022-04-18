package main

import (
	"Adesubomi/backend-2-challenge-golang/mods/maze"
	"Adesubomi/backend-2-challenge-golang/mods/user"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := *fiber.New()

	user.Bootstrap(&app)
	maze.Bootstrap(&app)

	app.Listen(":3080")
}
