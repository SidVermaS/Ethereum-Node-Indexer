package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	modules.InitializeAll()

	// Create an instance
	app := fiber.New()
	routes.SetupRoutes(app)
	
	PORT := fmt.Sprintf("%s", os.Getenv(string(consts.API_PORT)))

	log.Printf("Server is running on PORT: %s...\n", PORT)
	err := app.Listen(":" + PORT)
	if err != nil {
		panic(err)
	}

}
