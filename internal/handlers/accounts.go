package handlers

import (
	"log"
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

	// Convert string to uint64
	var accountID uint64
	accountID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println("Error parsing account id from URL to uint64")
		return c.SendStatus(fiber.StatusBadRequest)
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

func GetSelfAccount(c *fiber.Ctx) error {

	// Get account ID from JWT
	user := c.Locals("user").(*jwt.Token)
	if user == nil {
		log.Println("Error parsing user from claims")
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["id"] == nil {
		log.Println("Error parsing user id from claims")
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	accountID := claims["id"].(float64)
	if accountID == 0 {
		log.Println("Error parsing user id from claims")
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	accountIDUint64 := uint64(accountID)

	// Get account from database
	account := database.GetAccount(accountIDUint64)
	if account == nil {
		log.Println("Error getting account from database: not found")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Return account
	return c.JSON(account)
}
