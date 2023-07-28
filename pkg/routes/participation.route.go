package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)
// Setting up routes for indexing
func SetupParticipationRoutes(router fiber.Router) {
	participationRoutes := router.Group("/participation")
	// Fetch the Network's Participation Rate
	participationRoutes.Get("/network", controllers.GetNetworksParticipationRate)
	
	// Fetch an individual validator's Participation Rate
	participationRoutes.Get("/validators/:id", controllers.GetValidatorParticipationRate)
}
