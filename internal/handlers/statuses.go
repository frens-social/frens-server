package handlers

import (
	"log"
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type CreateStatusBody struct {
	Text string `json:"text" validate:"required"`
}

func CreateStatus(c *fiber.Ctx) error {

	// Parse request body
	var body CreateStatusBody
	if err := c.BodyParser(&body); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Parse account id from claims
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

	// Convert to status
	status := models.Status{
		Text:      body.Text,
		AccountID: accountIDUint64,
	}

	// Validate status
	if err := status.Validate(); err != nil {
		log.Println("Error validating status:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Insert status into database
	if ok := database.CreateStatus(&status); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return status
	return c.SendStatus(fiber.StatusOK)
}

func GetStatus(c *fiber.Ctx) error {

	// Get status id
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get status from database
	status, err := database.GetStatus(id)
	if err != nil {
		log.Println("Error getting status:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return status
	return c.JSON(status)
}

func GetStatuses(c *fiber.Ctx) error {

	// Read query parameters
	count := c.Query("count", "10")
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	// TODO: Add more query parameters

	// Get statuses from database
	statuses, err := database.GetStatuses(countInt)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return statuses
	return c.Status(fiber.StatusOK).JSON(statuses)
}

func DeleteStatus(c *fiber.Ctx) error {

	// Get status id
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Delete status from database
	if ok := database.DeleteStatus(id); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return status
	return c.SendStatus(fiber.StatusOK)
}
