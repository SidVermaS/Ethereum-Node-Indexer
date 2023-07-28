package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for epochs
func SetupEpochsRoutes(router fiber.Router) {
	epochsRoutes := router.Group("/epochs")
	
	// Fetch the paginated epochs
	epochsRoutes.Get("/", controllers.GetEpochs)
}
