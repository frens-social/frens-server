package handlers

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateAccountBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateAccount(c *fiber.Ctx) error {

	// Parse request body
	var body CreateAccountBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: validate body

	// TODO: Salt and hash password

	// Create account object
	var newAccount = models.Account{
		Username: body.Username,
	}

	// Insert account into database
	if ok := database.CreateAccount(&newAccount); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// TODO: API Response
	return c.SendStatus(fiber.StatusCreated)
}

func GetAllAccounts(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func GetAccount(c *fiber.Ctx) error {

	// Get account ID from URL
	id := c.Params("id", "")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get account from database
	var account models.Account
	if ok := database.GetAccount(id, &account); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return account
	return c.JSON(account)
}
