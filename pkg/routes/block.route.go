package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for blocks
func SetupBlocksRoutes(router fiber.Router) {
	blocksRoutes := router.Group("/blocks")
	
	// Fetch the paginated blocks
	blocksRoutes.Get("/", controllers.GetBlocks)
}
