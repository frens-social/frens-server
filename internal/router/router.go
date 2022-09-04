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
}

func addAuthenticatedRoutes() {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Authentication
	v1.Get("/auth/check", handlers.CheckAuth)

	// Accounts
	v1.Get("/accounts", handlers.GetAllAccounts)
	v1.Post("/accounts", handlers.CreateAccount)
	v1.Get("/accounts/:id", handlers.GetAccount)
	//v1.Get("/accounts/:id/statuses", handlers.GetAccountStatuses)
	//v1.Get("/accounts/:id/followers", handlers.GetAccountFollowers)
	//v1.Post("/accounts/:id/followers", handlers.FollowAccount)
	//v1.Get("/accounts/:id/following", handlers.GetAccountFollowing)

	// Reactions
	v1.Get("/statuses/:id/reactions", handlers.GetStatusReactions)
	v1.Get("statuses/:id/reactions/counts", handlers.GetStatusReactionsCounts)
	v1.Post("/statuses/:id/reactions", handlers.CreateReaction)

	// Statuses
	v1.Post("/statuses", handlers.CreateStatus)
	v1.Delete("/statuses/:id", handlers.DeleteStatus)
	v1.Get("/statuses/:id", handlers.GetStatus)

	// Feeds
	v1.Get("/feeds/home", handlers.GetHomeFeed)
	v1.Get("/feeds/public", handlers.GetPublicFeed)
	v1.Get("/feeds/federated", handlers.GetFederatedFeed)
	v1.Get("/feeds/user/self", handlers.GetSelfFeed)
	v1.Get("/feeds/user/:id", handlers.GetUserFeed)
	v1.Get("/feeds/tag/:tag", handlers.GetTagFeed)
}

func Run() {
	log.Fatal(app.Listen(":4000"))
}
