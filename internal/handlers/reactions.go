package handlers

import (
	"log"
	"strconv"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CreateReactionRequest struct {
	Emoji string `json:"emoji"`
}

func CreateReaction(c *fiber.Ctx) error {
	log.Println("fuck")

	// Parse the request body
	var req CreateReactionRequest
	if err := c.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return err
	}

	// Get user id
	userID, err := getRequestorID(c)
	if err != nil {
		log.Println("Error getting requestor id:", err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get status id
	statusID := c.Params("id")
	if statusID == "" {
		log.Println("Error getting status id:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	log.Println("Statusid:", statusID)
	statusIDuint, err := strconv.ParseUint(statusID, 10, 64)
	if err != nil {
		log.Println("Error parsing status id:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Create reaction
	reaction := models.Reaction{
		UserID:   *userID,
		StatusID: statusIDuint,
		Emoji:    req.Emoji,
	}

	// Insert reaction into database
	if err := database.CreateReaction(&reaction); err != nil {
		log.Println("Error creating reaction:", err)
		return err
	}

	// Return the reaction
	return c.Status(200).JSON(reaction)
}

func GetStatusReactions(c *fiber.Ctx) error {

	// Get status id
	statusID := c.Params("id")
	if statusID == "" {
		log.Println("Error getting status id")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Convert status id to uint
	statusIDuint, err := strconv.ParseUint(statusID, 10, 64)
	if err != nil {
		log.Println("Error parsing status id:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get query params
	userId := c.Query("user")

	// If userID is provided, convert to uint
	var userIDuint uint64
	if userId != "" {
		userIDuint, err = strconv.ParseUint(userId, 10, 64)
		if err != nil {
			log.Println("Error parsing user id:", err)
			return c.SendStatus(fiber.StatusBadRequest)
		}
	}

	// If userID is provided, get reactions for that user
	if userId != "" {
		reactions, err := database.GetUserStatusReactions(statusIDuint, userIDuint)
		if err != nil {
			log.Println("Error getting reactions:", err)
			return err
		}
		return c.Status(200).JSON(reactions)
	}

	// Otherwise, get all reactions
	reactions, err := database.GetStatusReactions(statusIDuint)
	if err != nil {
		log.Println("Error getting reactions:", err)
		return err
	}

	// Return the reactions
	return c.Status(200).JSON(reactions)
}

type ReactionCountEntry struct {
	Emoji string `json:"emoji"`
	Count int    `json:"count"`
}

func GetStatusReactionsCounts(c *fiber.Ctx) error {

	// Get status id
	statusID := c.Params("id")
	if statusID == "" {
		log.Println("Error getting status id")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Convert status id to uint
	statusIDuint, err := strconv.ParseUint(statusID, 10, 64)
	if err != nil {
		log.Println("Error parsing status id:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get reactions
	reactions, err := database.GetStatusReactions(statusIDuint)
	if err != nil {
		log.Println("Error getting reactions:", err)
		return err
	}

	// Count reactions
	reactionsCounts := []ReactionCountEntry{}
	for _, reaction := range reactions {
		found := false
		for i, reactionCount := range reactionsCounts {
			if reactionCount.Emoji == reaction.Emoji {
				reactionsCounts[i].Count++
				found = true
				break
			}
		}
		if !found {
			reactionsCounts = append(reactionsCounts, ReactionCountEntry{
				Emoji: reaction.Emoji,
				Count: 1,
			})
		}
	}

	// Return the reactions
	return c.Status(200).JSON(reactionsCounts)
}
