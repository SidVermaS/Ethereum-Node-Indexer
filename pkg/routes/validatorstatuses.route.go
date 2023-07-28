package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for validatorstatuses
func SetupValidatorStatusesRoutes(router fiber.Router) {
	validatorsRoutes := router.Group("/validator-statuses")
	
	// Fetch the paginated validatorstatuses
	validatorsRoutes.Get("/", controllers.GetValidatorStatuses)
}
