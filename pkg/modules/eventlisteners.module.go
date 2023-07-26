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
//	This handler can be used to implement complex error handling scenarios
func errorHandler(err error) error {
	log.Printf("~~~ EventListener error: %s", err)
	return err
}

var EpochsCount uint = 0
var AreEpochsSaved bool = false
var FinalizedCheckpoints []*consensysstructs.FinalizedCheckpoint

func ResetEpochsMetaData() {
	EpochsCount = 0
	AreEpochsSaved = false
}
func eventHandler(event *sseclient.Event) error {
	// Check the type of event for which the event was triggered.
	if event.Event == string(consensysconsts.Finalized_checkpoint) {
		// Only handles the 5 epochs incoming in every 12.40 minutes.
		if EpochsCount < 5 {
			log.Printf("~~~ event : %s : %s : %s ", event.ID, event.Event, event.Data)
			// Increment the EpochsCount
			EpochsCount = EpochsCount + 1
			var finalizedCheckpoint *consensysstructs.FinalizedCheckpoint
			// Convert from bytes to a struct
			err := json.Unmarshal(event.Data, &finalizedCheckpoint)
			// Returns an error 
			if err != nil {
				return err
			}
			// Appending the incoming epoch data needed for further processing and indexing
			FinalizedCheckpoints = append(FinalizedCheckpoints, finalizedCheckpoint)
		} else if !AreEpochsSaved {
			AreEpochsSaved = true
			// If 5 epochs are fetched then we save the data in our database
			ProcessToSaveDataForIndexing(FinalizedCheckpoints)
		}
	}
	var err error = nil
	return err
}
func StreamConsensysNode(consensysVendor *consensys.Consensys, topicsSlice []consensysconsts.ConsensysTopicsE) {
	// List of topics to which we have to subscribe
	var topicsStringSlice []string = consensyshelpers.ConvertTopicsSliceToStringSlice(topicsSlice)
	//  Converting from slice to a string
	var topics string = strings.Join(topicsStringSlice, ",")
	// URL for listening to the events
	var u string = fmt.Sprintf("%s/eth/v1/events?topics=%s", consensysVendor.Vendor.BaseURL, topics)
	// It creates SSE stream client object
	eventSource := sseclient.New(u, "")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// Connects to the SSE stream
	eventSource.Start(ctx, eventHandler, errorHandler)
	// Canceling this context releases resources associated with it
	defer cancel()
}
