package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to the DB, cache, start node scheduler and start event listenter
	modules.ActivateAll()

	// Create an instance of the fiber app
	app := fiber.New()
	// Enable cors
	app.Use(cors.New())
	// Setting up the API routes
	routes.SetupRoutes(app)

	// Fetching the PORT from the environment variable
	PORT := fmt.Sprintf("%s", os.Getenv(string(consts.API_PORT)))

	log.Printf("Server is running on PORT: %s...\n", PORT)

	// Listening to requests on the PORT
	err := app.Listen(":" + PORT)
	if err != nil {
		panic(err)
	}

}
