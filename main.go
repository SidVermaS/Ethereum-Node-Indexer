package main

import (
	"github.com/SidVermaS/Ethereum-Consensus-Layer/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func main()	{
	config.InitializeAll()
	app:=fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello ECL!")
	})

	app.Listen(":8080")
}