package handlers

import (
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateStatusBody struct {
	Text string `json:"text" validate:"required"`
}

func CreateStatus(c *fiber.Ctx) error {

	// Parse request body
	var body CreateStatusBody
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Convert to status
	status := models.Status{
		Text: body.Text,
	}

	// Insert status into database
	if ok := database.CreateStatus(&status); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return status
	return c.Status(fiber.StatusOK).JSON(status)
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

func GetStatus(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
