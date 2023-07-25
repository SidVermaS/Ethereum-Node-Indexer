package controllers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GetValidators(ctx *fiber.Ctx) error {
	offset, limit := helpers.GetPaginationValues(ctx)
	validators, err := modules.GetValidators(offset, limit)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"validators": validators})
	return nil
}
