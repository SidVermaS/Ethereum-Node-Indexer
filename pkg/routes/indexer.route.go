package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupIndexerRoutes(router fiber.Router) {
	indexerRoutes := router.Group("/indexers")
	indexerRoutes.Get("/", controllers.GetNetworksParticipationRate)
	// indexerRoutes.Get("/:id", controllers.GetValidatorsParticipationRate)
}
