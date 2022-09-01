package handlers

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

func GetHomeFeed(c *fiber.Ctx) error {

	// PLACEHOLDER

	// Get statuses from database
	statuses, err := database.GetPublicFeed(nil)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Send statuses to client
	return c.JSON(statuses)

}

func GetPublicFeed(c *fiber.Ctx) error {

	// Params
	// Placeholder for now. This will be used to requery when scrolling.
	//continue := c.Query("continue")

	// Get statuses from database
	statuses, err := database.GetPublicFeed(nil)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Send statuses to client
	return c.JSON(statuses)
}

func GetFederatedFeed(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func GetUserFeed(c *fiber.Ctx) error {

	// Params
	userId := c.Params("id")

	// Get statuses from database
	statuses, err := database.GetUserFeed(userId, nil)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Send statuses to client
	return c.JSON(statuses)
}

func GetSelfFeed(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func GetTagFeed(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
