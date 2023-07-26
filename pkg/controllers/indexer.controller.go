package controllers

import (
	"fmt"

	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

func GetNetworksParticipationRate(ctx *fiber.Ctx) error {
	cachedEpochId, _ := configs.GetCacheValue(consts.LAST_EPOCH_SAVED)
	var needToFetchNetworksParticipationRate bool = true
	var networksParticipationRate float64
	if cachedEpochId != "" {
		epoch, err := modules.GetLastEpoch()
		if err == nil {
			epochId, _ := helpers.ConvertStringToUInt(cachedEpochId)
			if epoch.ID == epochId {
				cachedNetworksParticipationRate, err := configs.GetCacheValue(consts.NETWORKS_PARTICIPATION_RATE)
				if err == nil && cachedNetworksParticipationRate != "" {
					tempNetworksParticipationRate, err := helpers.ConvertStringToFloat(cachedNetworksParticipationRate)
					if err == nil {
						needToFetchNetworksParticipationRate = false
						networksParticipationRate = tempNetworksParticipationRate
					}
				}
			}
		}
	}
	if needToFetchNetworksParticipationRate {
		epoch, _ := modules.GetLastEpoch()
		tempNetworksParticipationRate, err := modules.GetNetworksParticipationRate()
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
			return err
		}
		networksParticipationRate = tempNetworksParticipationRate
		configs.SetCacheValue(consts.LAST_EPOCH_SAVED, epoch.ID)
		configs.SetCacheValue(consts.NETWORKS_PARTICIPATION_RATE, networksParticipationRate)
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
