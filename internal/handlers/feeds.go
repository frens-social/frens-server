package handlers

import (
	"log"

	"github.com/bwoff11/frens/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetHomeFeed(c *fiber.Ctx) error {

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

	// Get list of followed accounts from database
	followedAccounts, err := database.GetFollowedAccounts(accountIDUint64)

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
