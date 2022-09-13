package router

import (
	"log"

	"github.com/bwoff11/frens/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

var app *fiber.App

func Create() {
	app = fiber.New()

	addMiddleware()
	// TODO: Get config from config package

	addUnauthenticatedRoutes()

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	}))

	addAuthenticatedRoutes()
}

func addMiddleware() {
	app.Use(cors.New())
	app.Use(logger.New())
}

func addUnauthenticatedRoutes() {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Authentication
	v1.Post("/auth/login", handlers.Login)

	// Media
	v1.Get("/media/:id", handlers.GetMedia)
}

func addAuthenticatedRoutes() {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Authentication
	v1.Get("/auth/check", handlers.CheckAuth)

	// Users
	v1.Get("/users", handlers.GetAllUsers)
	v1.Post("/users", handlers.CreateUser)
	v1.Get("/users/:id", handlers.GetUser)
	//v1.Get("/users/:id/statuses", handlers.GetUserStatuses)
	//v1.Get("/users/:id/followers", handlers.GetUserFollowers)
	//v1.Post("/users/:id/followers", handlers.FollowUser)
	//v1.Get("/users/:id/following", handlers.GetUserFollowing)

	// Status Media
	v1.Get("/statuses/:statusID/media", handlers.GetStatusMedia)
	v1.Post("/statuses/:statusID/media", handlers.CreateStatusMedia)
	v1.Delete("/statuses/:statusID/media/:mediaID", handlers.DeleteStatusMedia)

	// Reactions
	v1.Get("/statuses/:id/reactions", handlers.GetStatusReactions)
	v1.Get("statuses/:id/reactions/counts", handlers.GetStatusReactionsCounts)
	v1.Post("/statuses/:id/reactions", handlers.CreateReaction)

	// Statuses
	v1.Get("/statuses/:id", handlers.GetStatus)
	v1.Patch("/statuses/:id", handlers.UpdateStatus)
	v1.Delete("/statuses/:id", handlers.DeleteStatus)
	v1.Post("/statuses", handlers.CreateStatus)

	// Feeds
	v1.Get("/feeds/home", handlers.GetHomeFeed)
	v1.Get("/feeds/public", handlers.GetPublicFeed)
	v1.Get("/feeds/federated", handlers.GetFederatedFeed)
	v1.Get("/feeds/user/self", handlers.GetSelfFeed)
	v1.Get("/feeds/user/:id", handlers.GetUserFeed)
	v1.Get("/feeds/tag/:tag", handlers.GetTagFeed)

	// Media
	v1.Post("/media", handlers.UploadMedia)
}

func Run() {
	log.Fatal(app.Listen(":4000"))
}
