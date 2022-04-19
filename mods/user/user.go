package user

import (
	"Adesubomi/backend-2-challenge-golang/pkg"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Username *string `gorm:"primaryKey;" json:"username"`
	Password *string `gorm:"type:varchar(191);not null" json:"password"`
}

func bootDatabase() *gorm.DB {
	// setup the database
	config := pkg.DatabaseConfig{
		Host:     "localhost",
		Port:     "3306",
		DBName:   "maze-runner",
		Username: "root",
		Password: "",
	}

	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	db.DB()
	return db
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

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(ud.Password), 14)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"message": "Unable to hash password"},
		)
	}

	ud.Password = string(passwordBytes)

	result := db.Create(&ud)

	if result.Error != nil {
		return c.Status(500).JSON(
			fiber.Map{"message": "Unable to create account"},
		)
	}

	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "User account has been created"},
	)
}

func Bootstrap(f *fiber.App) {
	db = bootDatabase()

	f.Post("/login", login)
	f.Post("/user", registerUser)
}
