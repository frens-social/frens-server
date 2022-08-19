package router

import (
	"log"

	"github.com/bwoff11/frens/internal/handlers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var app *fiber.App

func Create() {
	app = fiber.New()

	// TODO: Add middleware
	// TODO: Get config from config package

	addUnauthenticatedRoutes()

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	addAuthenticatedRoutes()
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

	// Accounts
	v1.Get("/accounts", handlers.GetAllAccounts)
	v1.Post("/accounts", handlers.CreateAccount)
	v1.Get("/accounts/:id", handlers.GetAccount)

	// Statuses
	v1.Post("/statuses", handlers.CreateStatus)
	v1.Delete("/statuses/:id", handlers.DeleteStatus)
	v1.Get("/statuses/:id", handlers.GetStatus)

	// Favourites
	//v1.Post("/status/:id/favourites", handlers.GetFavouritedBy)
	//v1.Post("/status/:id/favourites", handlers.CreateFavourite)
	//v1.Delete("/status/:id/favourites", handlers.DeleteFavourite)

	// Feeds

}

func Run() {
	log.Fatal(app.Listen(":4000"))
}
