package helpers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	sseclient "github.com/advbet/sseclient"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys"
	consensysconsts "github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys/consts"
	consensysutils "github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys/utils"
)

func errorHandler(err error) error {
	log.Printf("~~~ EventListener error: %s", err)
	return err
}

var Count uint = 0

func eventHandler(event *sseclient.Event) error {
	if Count < 5 {
		log.Printf("~~~ event : %s : %s : %s ", event.ID, event.Event, event.Data)
		Count = Count + 1
	}
	var err error = nil
	return err
}
func StreamConsensysNode(consensysVendor *consensys.Consensys, topicsSlice []consensysconsts.ConsensysTopicsE) {

	var topicsStringSlice []string = consensysutils.ConvertTopicsSliceToStringSlice(topicsSlice)
	//	http://localhost:5051
	var topics string = strings.Join(topicsStringSlice, ",")
	var u string = fmt.Sprintf("%s/eth/v1/events?topics=%s", consensysVendor.Vendor.BaseURL, topics)
	eventSource := sseclient.New(u, "")
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()
	eventSource.Start(ctx, eventHandler, errorHandler)

}
