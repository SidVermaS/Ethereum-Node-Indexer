package controllers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GetEpochs(ctx *fiber.Ctx) error {
	// Get values for the pagination
	offset, limit := helpers.GetPaginationValues(ctx)

	// Queries returns a map of query parameters and their values.
	queries := ctx.Queries()

	// Fetches the paginated epochs
	epochs, err := modules.GetEpochs(queries, offset, limit)

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	// Returns an array of the epochs in JSON format
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"epochs": epochs})
	return nil
}
