package handlers

import (
	"log"
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(c *fiber.Ctx) error {

	// Parse request body
	var body CreateUserBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: validate body

	// TODO: Salt and hash password

	// Create User object
	var newUser = models.User{
		Username: body.Username,
	}

	// Insert User into database
	if ok := database.CreateUser(&newUser); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// TODO: API Response
	return c.SendStatus(fiber.StatusCreated)
}

func GetAllUsers(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func GetUser(c *fiber.Ctx) error {

	// Get User ID from URL
	id := c.Params("id", "")
	if id == "" {
		log.Println("Error parsing User id from URL")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get User ID
	var UserID *uint64
	// If ID is self, get User from JWT
	if id == "self" {
		var err error
		UserID, err = getRequestorID(c)
		if err != nil {
			log.Println("Error getting requestor ID from JWT")
			return err
		}
	} else {
		parsedUserID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Println("Error parsing User id from URL to uint64")
			return c.SendStatus(fiber.StatusBadRequest)
		}
		UserID = &parsedUserID
	}

	// Get User from database
	User := database.GetUser(UserID)
	if User == nil {
		log.Println("Error getting User from database: not found")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Return User
	return c.JSON(User)
}
