package main

import (
	"Adesubomi/backend-2-challenge-golang/mods/maze"
	"Adesubomi/backend-2-challenge-golang/mods/user"
	"Adesubomi/backend-2-challenge-golang/pkg"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := *fiber.New(*pkg.RouterConfig())

	user.Bootstrap(&app)
	maze.Bootstrap(&app)

	log.Fatal(
		app.Listen(":4000"),
	)
}
