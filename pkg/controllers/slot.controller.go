package controllers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GetSlots(ctx *fiber.Ctx) error {
	// Get values for the pagination
	offset, limit := helpers.GetPaginationValues(ctx)

	// Queries returns a map of query parameters and their values.
	queries := ctx.Queries()

	// Fetches the filtered paginated slots
	slots, err := modules.GetFilteredSlots(queries, offset, limit)

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	// Returns an array of the slots in JSON format
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"slots": slots})
	return nil
}
