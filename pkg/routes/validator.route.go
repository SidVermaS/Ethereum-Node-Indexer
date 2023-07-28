package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for validators
func SetupValidatorsRoutes(router fiber.Router) {
	validatorsRoutes := router.Group("/validators")
	
	// Fetch the paginated validators
	validatorsRoutes.Get("/", controllers.GetValidators)
}
