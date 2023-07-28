package controllers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GetBlocks(ctx *fiber.Ctx) error {
	// Get values for the pagination
	offset, limit := helpers.GetPaginationValues(ctx)

	// Fetches the paginated blocks
	blocks, err := modules.GetBlocks(offset, limit)

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	// Returns an array of the blocks in JSON format
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"blocks": blocks})
	return nil
}
