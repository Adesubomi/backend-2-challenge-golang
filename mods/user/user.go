package user

import (
	"Adesubomi/backend-2-challenge-golang/pkg"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"type:varchar(191);not null" json:"password"`
}

func registerUser(c *fiber.Ctx) error {
	var uData User
	if err := c.BodyParser(&uData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Unable to create new user",
				"error":   err,
			},
		)
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(string(uData.Password)), 14)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{"message": "Unable to hash password"},
		)
	}

	uData.Password = string(passwordBytes)

	result := db.Create(&uData)

	if result.Error != nil {
		return c.Status(500).JSON(
			fiber.Map{"message": "Unable to create account"},
		)
	}

	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "User account has been created"},
	)
}

func getUserByUsername(n string) (*User, error) {
	var user User

	if err := db.Where(&User{Username: n}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func Bootstrap(f *fiber.App) {
	db = pkg.GetDatabaseConnection(User{})
	f.Post("/login", login)
	f.Post("/user", registerUser)
}
