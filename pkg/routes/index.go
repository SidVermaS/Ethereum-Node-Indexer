package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	v1Routes := app.Group("/api/v1")

	SetupIndexerRoutes(v1Routes)
	SetupValidatorsRoutes(v1Routes)
}
