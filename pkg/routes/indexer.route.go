package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)
// Setting up routes for indexing
func SetupIndexerRoutes(router fiber.Router) {
	indexerRoutes := router.Group("/indexers")
	// Fetch the Network's Participation Rate
	indexerRoutes.Get("/network", controllers.GetNetworksParticipationRate)
	
	// Fetch an individual validator's Participation Rate
	indexerRoutes.Get("/validators/:id", controllers.GetValidatorParticipationRate)
}
