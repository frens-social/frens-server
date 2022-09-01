package handlers

import (
	"log"
	"strconv"

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
		log.Println("Error parsing account id from URL")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get account ID
	var accountID *uint64
	// If ID is self, get account from JWT
	if id == "self" {
		var err error
		accountID, err = getRequestorID(c)
		if err != nil {
			log.Println("Error getting requestor ID from JWT")
			return err
		}
	} else {
		parsedAccountID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Println("Error parsing account id from URL to uint64")
			return c.SendStatus(fiber.StatusBadRequest)
		}
		accountID = &parsedAccountID
	}

	// Get account from database
	account := database.GetAccount(accountID)
	if account == nil {
		log.Println("Error getting account from database: not found")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Return account
	return c.JSON(account)
}
