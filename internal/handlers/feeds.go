package handlers

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
)

func GetHomeFeed(c *fiber.Ctx) error {

	/*
		// Get account ID from JWT
		accountID, err := getRequestorID(c)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Get list of followed accounts from database
		followedAccounts, err := database.GetFollowedAccountIDs(accountID)

		// Include self in list of followed accounts
		followedAccounts = append(followedAccounts, *accountID)

	*/

	return c.SendStatus(fiber.StatusNotImplemented)

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
	return c.SendStatus(fiber.StatusNotImplemented)
}

func GetSelfFeed(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func GetTagFeed(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
