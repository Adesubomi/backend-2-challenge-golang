package user

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
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
			fiber.Map{"message": "UserModel not found"},
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

	claims := jwt.MapClaims{
		"user":    user,
		"expires": time.Now().AddDate(0, 0, 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a
	// string using the secret, without a secret
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

func Protected(c *fiber.Ctx) error {

	authorization := c.GetReqHeaders()["Authorization"]
	tParts := strings.Split(authorization, " ")

	if len(tParts) != 2 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{})
	}

	tokenString := tParts[1]
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Token",
			"error":   err.Error(),
		})
	}

	expiryTime := claims["expires"].(time.Time)

	if expiryTime.Before(time.Now()) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token has expired",
		})
	}

	jsonBody, err := json.Marshal(claims["user"])

	var user *User
	if err := json.Unmarshal(jsonBody, &user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to process your request. Data is malformed",
			"error":   err.Error(),
		})
	}

	c.Locals("user", user)

	return c.Next()
}
