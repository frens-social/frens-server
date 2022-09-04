package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func getRequestorID(c *fiber.Ctx) (*uint64, error) {
	user := c.Locals("user").(*jwt.Token)
	if user == nil {
		log.Println("Error parsing user from claims")
		return nil, c.SendStatus(fiber.StatusUnauthorized)
	}
	claims := user.Claims.(jwt.MapClaims)
	if claims["id"] == nil {
		log.Println("Error parsing user id from claims")
		return nil, c.SendStatus(fiber.StatusUnauthorized)
	}
	userID := claims["id"].(float64)
	if userID == 0 {
		log.Println("Error parsing user id from claims")
		return nil, c.SendStatus(fiber.StatusUnauthorized)
	}

	uint64ID := uint64(userID)
	return &uint64ID, nil
}
