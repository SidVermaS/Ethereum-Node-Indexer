package consensys

import (
	"encoding/json"
	"fmt"
	"log"

	vstructs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendorpkg/structs"

	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys/consensysstructs"
)

type Consensys struct {
	Vendor vstructs.Vendor
	Topics []string
}
// Fetches validators based on the stateId
func (consensys *Consensys) GetValidatorsFromState(stateIdentifierOrHex string) *consensysstructs.GetValidatorsFromStateResponse {
	_, response, _ := consensys.Vendor.CallAPI(&vstructs.APIRequest{Url: fmt.Sprintf("/eth/v1/beacon/states/%s/validators", stateIdentifierOrHex)})
	var getValidatorsFromStateResponse *consensysstructs.GetValidatorsFromStateResponse
	err := json.Unmarshal(response, &getValidatorsFromStateResponse)

	if err != nil {
		log.Printf("~~~ Error GetValidatorsFromState() %s\n", err)
	}
	return getValidatorsFromStateResponse
}

// Fetches committees based on the stateId and epoch
func (consensys *Consensys) GetCommitteesAtState(stateIdentifierOrHex string, epoch uint) *consensysstructs.GetCommitteesAtStateResponse {
	_, response, _ := consensys.Vendor.CallAPI(&vstructs.APIRequest{Url: fmt.Sprintf("/eth/v1/beacon/states/%s/committees?epoch=%d", stateIdentifierOrHex, epoch)})
	var getCommitteesAtStateResponse *consensysstructs.GetCommitteesAtStateResponse
	err := json.Unmarshal(response, &getCommitteesAtStateResponse)

	if err != nil {
		log.Printf("~~~ Error GetCommitteesAtState() %s\n", err)
	}
	return getCommitteesAtStateResponse
}

