package maze

import (
	"Adesubomi/backend-2-challenge-golang/mods/user"
	"Adesubomi/backend-2-challenge-golang/pkg"
	"fmt"
	"gorm.io/gorm"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var db *gorm.DB

type Maze struct {
	gorm.Model
	UserID   int
	Label    string    `gorm:"unique" json:"label"`
	GridSize string    `gorm:"" json:"grid_size"`
	Walls    string    `gorm:"" json:"walls"`
	Entrance string    `gorm:"" json:"entrance"`
	User     user.User `gorm:"foreignKey:UserID"`
}

func getMazes(c *fiber.Ctx) error {

	u := c.Locals("user").(user.User)

	fmt.Println(u.Username)

	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "List of all Mazes"},
	)
}

func storeMaze(c *fiber.Ctx) error {

	var mazeData Maze

	if err := c.BodyParser(&mazeData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to create new User",
		})
	}

	result := db.Create(&mazeData)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed creating new maze",
			"error":   result.Error,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User has been created.",
		"data":    result.Row(),
	})
}

func getMazeSolution(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Shortest path of User"},
	)
}

func Bootstrap(f *fiber.App) {
	db = pkg.GetDatabaseConnection(Maze{})

	f.Post("/maze", user.Protected, storeMaze)
	f.Get("/maze", user.Protected, getMazes)
	f.Get("/maze/:id/solution", user.Protected, getMazeSolution)
}
