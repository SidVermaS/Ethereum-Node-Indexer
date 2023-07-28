package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setting up routes for states
func SetupStatesRoutes(router fiber.Router) {
	statesRoutes := router.Group("states")
	
	// Fetch the paginated states
	statesRoutes.Get("/", controllers.GetStates)
}
