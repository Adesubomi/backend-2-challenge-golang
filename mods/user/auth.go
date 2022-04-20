package user

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c *fiber.Ctx) error {

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Login request is invalid"},
		)
	}

	user, err := getUserByUsername(input.Username)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Unable to verify user account!"},
		)
	} else if user == nil {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "User not found"},
		)
	}

	dbHashedPassword := user.Password
	stringPassword := input.Password

	matchTest := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(stringPassword)) == nil

	if !matchTest {
		return c.Status(http.StatusUnauthorized).JSON(
			fiber.Map{"message": "Invalid Credentials"},
		)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a
	// string using the secret, without a seccret
	tokenString, err := token.SignedString([]byte{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{"message": "Unable to log you in at the moment."},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "Login Successful!",
			"data": fiber.Map{
				"token": tokenString,
			},
		},
	)
}
