package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadMedia(c *fiber.Ctx) error {

	// Get file
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error getting file from form: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Assign random uuid
	id := uuid.New().String()

	// Save file
	err = c.SaveFile(file, "./media/"+id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Return success
	return c.JSON(fiber.Map{
		"success": true,
		"id":      id,
	})
}

func GetMedia(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendFile("./media/" + id)
}
