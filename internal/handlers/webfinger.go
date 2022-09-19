package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ResourceType string

const (
	AccountResourceType ResourceType = "acct"
)

type WebfingerLink struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type WebfingerResponse struct {
	Subject string          `json:"subject"`
	Aliases []string        `json:"aliases"`
	Links   []WebfingerLink `json:"links"`
}

func GetWebfinger(c *fiber.Ctx) error {

	// Get the resource from the query string
	resourceQuery := c.Query("resource")
	if resourceQuery == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "No resource provided",
		})
	}

	// Split the query string into the resource type and the resource
	resourceParts := strings.Split(resourceQuery, ":")
	if len(resourceParts) != 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid resource provided",
		})
	}
	resourceType := resourceParts[0]
	resource := resourceParts[1]

	// Switch on resource type
	switch ResourceType(resourceType) {
	case AccountResourceType:
		return getAccountWebfinger(c, resource)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid resource type",
		})
	}
}

func getAccountWebfinger(c *fiber.Ctx, resource string) error {
	resp := WebfingerResponse{
		Subject: "acct:admin@localhost",
		Aliases: []string{
			"http://localhost:3001/users/autumnfire",
		},
		Links: []WebfingerLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: "http://localhost:3001/users/autumnfire",
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: "http://localhost:3001/users/autumnfire",
			},
		},
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
