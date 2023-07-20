package main

import (
	
	"github.com/SidVermaS/Ethereum-Consensus/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitializeAll()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello EC!")
	})

	app.Listen(":8080")
}
