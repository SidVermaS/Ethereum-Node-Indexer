package routes

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupValidatorsRoutes(router fiber.Router) {
	validatorsRoutes := router.Group("/validators")
	validatorsRoutes.Get("/", controllers.GetValidators)
}
