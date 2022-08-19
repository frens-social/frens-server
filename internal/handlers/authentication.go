package handlers

import (
	"time"

	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {

	// Parse request body
	var body LoginBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Authenticate user
	if ok := database.Authenticate(body.Username, body.Password); !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	})
}
