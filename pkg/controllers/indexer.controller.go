package controllers

import (
	"fmt"

	configs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/config"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/helpers"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/modules"
	"github.com/gofiber/fiber/v2"
)

// Fetch the network's participation rate
func GetNetworksParticipationRate(ctx *fiber.Ctx) error {
	// Fetching the epoch id from the Redis Cache
	cachedEpochId, _ := configs.GetCacheValue(consts.LAST_EPOCH_SAVED)
	var needToFetchNetworksParticipationRate bool = true
	var networksParticipationRate float64
	// Initially, we check whether we have the participation rate in our cache or not, if it's availabe in the cache then we fetch it from there, otherwise, we fetch it from the DB and save in the cache for the next requests
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
		epoch, err := modules.GetLastEpoch()
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Please wait for 4 minutes, we are processing the data..."})
			return err
		}
		// Fetches from the DB
		tempNetworksParticipationRate, err := modules.GetNetworksParticipationRate()
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
			return err
		}
		networksParticipationRate = tempNetworksParticipationRate
		// Saves it in cache
		configs.SetCacheValue(consts.LAST_EPOCH_SAVED, epoch.ID)
		configs.SetCacheValue(consts.NETWORKS_PARTICIPATION_RATE, networksParticipationRate)
	}
	// Returns the network's participation rate in JSON format
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"participation_rate": fmt.Sprintf("%.2f%%", networksParticipationRate)})
	return nil
}

// Fetch an individual validator's Participation Rate
func GetValidatorParticipationRate(ctx *fiber.Ctx) error {
	// Gets the validator's id
	idParam := ctx.Params("id")
	id, err := helpers.ConvertStringToUInt(idParam)
	
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		return nil
	}
	networksParticipationRate, err := modules.GetValidatorsParticipationRate(id)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong..."})
		return err
	}
	// Returns an individual validator's participation rate in JSON format
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"participation_rate": fmt.Sprintf("%.2f%%", networksParticipationRate)})
	return nil
}
