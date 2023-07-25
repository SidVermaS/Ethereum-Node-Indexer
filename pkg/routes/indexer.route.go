package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupIndexerRoutes(router fiber.Router) {
	indexerRoutes := router.Group("/indexers")
	indexerRoutes.Get("/network", controllers.GetNetworksParticipationRate)
	indexerRoutes.Get("/validators/:id", controllers.GetValidatorParticipationRate)
}
