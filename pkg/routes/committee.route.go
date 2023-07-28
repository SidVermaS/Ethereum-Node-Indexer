package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for committees
func SetupCommitteeRoutes(router fiber.Router) {
	committeesRoutes := router.Group("/committees")
	
	// Fetch the paginated committees
	committeesRoutes.Get("/", controllers.GetCommittees)
}
