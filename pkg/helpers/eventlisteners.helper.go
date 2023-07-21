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
	log.Printf("error: %s", err)
	return err
}
func eventHandler(event *sseclient.Event) error {
	log.Printf("event : %s : %s : %d bytes", event.ID, event.Event, len(event.Data))
	var err error = nil
	return err
}
func StreamConsensysNode(consensysVendor *consensys.Consensys, topicsSlice []consensysconsts.ConsensysTopicsE) {

	var topicsStringSlice []string = consensysutils.ConvertTopicsSliceToStringSlice(topicsSlice)

	var topics string = strings.Join(topicsStringSlice, ",")
	topics = topics[0 : len(topics)-1]
	// consensys.Consensys.Vendor
	var u string = fmt.Sprintf("%s?topics=%s", consensysVendor.Vendor.BaseURL, "")
	log.Printf("~~~ url: %s\n", u)
	eventSource := sseclient.New(u, "")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	eventSource.Start(ctx, eventHandler, errorHandler)

}
