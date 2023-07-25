package controllers

import (
	"fmt"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GetNetworksParticipationRate(ctx *fiber.Ctx) error {
	networksParticipationRate, err := modules.GetNetworksParticipationRate()
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"participation_rate": fmt.Sprintf("%.2f%%", networksParticipationRate)})
	return nil
}

func GetValidatorParticipationRate(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := helpers.ConvertStringToUInt(idParam)
	networksParticipationRate, err := modules.GetValidatorsParticipationRate(id)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"participation_rate": fmt.Sprintf("%.2f%%", networksParticipationRate)})
	return nil
}
