package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for slots
func SetupSlotsRoutes(router fiber.Router) {
	slotsRoutes := router.Group("/slots")
	
	// Fetch the paginated slots
	slotsRoutes.Get("/", controllers.GetSlots)
}
