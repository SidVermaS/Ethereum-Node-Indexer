package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	v1Routes := app.Group("/api/v1")
	// Routes for indexing the data
	SetupIndexerRoutes(v1Routes)
	// Routes for the validators' data
	SetupValidatorsRoutes(v1Routes)
}
