package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	sseclient "github.com/advbet/sseclient"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys"
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consensysstructs"
	consensysconsts "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consts"
	consensyshelpers "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/helpers"
)

func errorHandler(err error) error {
	log.Printf("~~~ EventListener error: %s", err)
	return err
}

var EpochsCount uint = 0
var AreEpochsSaved bool = false
var FinalizedCheckpoints []*consensysstructs.FinalizedCheckpoint

func ResetEpochsMetaData()	{
	EpochsCount=0
	AreEpochsSaved=false
}
func eventHandler(event *sseclient.Event) error {
	if event.Event == string(consensysconsts.Finalized_checkpoint) {
		if EpochsCount < 5 {
			log.Printf("~~~ event : %s : %s : %s ", event.ID, event.Event, event.Data)
			EpochsCount = EpochsCount + 1
			var finalizedCheckpoint *consensysstructs.FinalizedCheckpoint
			err := json.Unmarshal(event.Data, &finalizedCheckpoint)
			if err != nil {
				return err
			}
			FinalizedCheckpoints = append(FinalizedCheckpoints, finalizedCheckpoint)
		} else if !AreEpochsSaved {
			AreEpochsSaved = true
			ProcessToSaveDataForIndexing(FinalizedCheckpoints)
		}
	}
	var err error = nil
	return err
}
func StreamConsensysNode(consensysVendor *consensys.Consensys, topicsSlice []consensysconsts.ConsensysTopicsE) {

	var topicsStringSlice []string = consensyshelpers.ConvertTopicsSliceToStringSlice(topicsSlice)
	var topics string = strings.Join(topicsStringSlice, ",")
	var u string = fmt.Sprintf("%s/eth/v1/events?topics=%s", consensysVendor.Vendor.BaseURL, topics)

	eventSource := sseclient.New(u, "")
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()
	eventSource.Start(ctx, eventHandler, errorHandler)

}
