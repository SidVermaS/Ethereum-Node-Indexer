package routes

import "github.com/gofiber/fiber/v2"

// Setup all routes in the app
func SetupRoutes(app *fiber.App) {
	v1Routes := app.Group("/api/v1")
	// Routes for participation rate's data
	SetupParticipationRoutes(v1Routes)
	// Routes for the blocks' data
	SetupBlocksRoutes(v1Routes)
	// Routes for the committee' data
	SetupCommitteeRoutes(v1Routes)
	// Routes for the epochs' data
	SetupEpochsRoutes(v1Routes)
	// Routes for the slots' data
	SetupSlotsRoutes(v1Routes)
	// Routes for the states' data
	SetupStatesRoutes(v1Routes)
	// Routes for the validators' data
	SetupValidatorsRoutes(v1Routes)
	// Routes for the validator statuses' data
	SetupValidatorStatusesRoutes(v1Routes)
}
