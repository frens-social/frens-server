package handlers

import (
	"log"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateStatusBody struct {
	Text     string                `json:"text" validate:"required"`
	MediaIDs []string              `json:"media_ids"`
	Privacy  *models.StatusPrivacy `json:"privacy" validate:"required"`
	Draft    *bool                 `json:"draft" validate:"required"`
}

type UpdateStatusRequest struct {
	Text    string               `json:"text" validate:"required"`
	Privacy models.StatusPrivacy `json:"privacy" validate:"required"`
}

func CreateStatus(c *fiber.Ctx) error {

	// Parse request body
	var body CreateStatusBody
	if err := c.BodyParser(&body); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get user id
	userID, err := getRequestorID(c)
	if err != nil || userID == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Convert to status
	status := models.Status{
		Text:    body.Text,
		UserID:  *userID,
		Privacy: models.StatusPrivacyPublic, // Placeholder, need to get optional value from body
		Draft:   false,                      // Placeholder, need to get optional value from body
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

	// Convert media ids to uuids
	var mediaIDs []uuid.UUID
	for _, mediaID := range body.MediaIDs {
		uuid, err := uuid.Parse(mediaID)
		if err != nil {
			log.Println("Error parsing media id:", err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
		mediaIDs = append(mediaIDs, uuid)
	}

	// Insert media into database
	for _, mediaID := range mediaIDs {
		statusMedia := models.StatusMedia{
			ID:       mediaID,
			StatusID: status.ID,
		}
		if ok := database.CreateStatusMedia(&statusMedia); !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	// Return status
	return c.Status(fiber.StatusCreated).JSON(status)
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

	// Delete StatusMedia from database
	// TODO: This should be replaced with cascading deletes
	if ok := database.DeleteStatusMediaByStatusID(id); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// TODO: Add file deletion

	// Return status
	return c.SendStatus(fiber.StatusOK)
}

func UpdateStatus(c *fiber.Ctx) error {

	// Get status id
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Parse request body
	var req UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get user id
	userID, err := getRequestorID(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get original status
	status, err := database.GetStatus(id)
	if err != nil {
		log.Println("Error getting status:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Check if user is authorized to update status
	if status.UserID != *userID {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Update status
	if req.Text != "" {
		status.Text = req.Text
	}
	if req.Privacy != "" {
		status.Privacy = req.Privacy
	}

	// Validate status
	if err := status.Validate(); err != nil {
		log.Println("Error validating status:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Update status in database
	if ok := database.UpdateStatus(status); !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Return status
	return c.SendStatus(fiber.StatusOK)
}
