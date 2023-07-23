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

func (consensys *Consensys) GetValidatorsFromState(stateIdentifierOrHex string) *consensysstructs.GetValidatorsFromStateResponse {
	_, response, _ := consensys.Vendor.CallAPI(&vstructs.APIRequest{Url: fmt.Sprintf("/eth/v1/beacon/states/%s/validators", stateIdentifierOrHex)})
	var getValidatorsFromStateResponse *consensysstructs.GetValidatorsFromStateResponse
	err := json.Unmarshal(response, &getValidatorsFromStateResponse)
	
	
	if err != nil {
		log.Printf("~~~ Error GetValidatorsFromState() %s\n", err)
	}
	return getValidatorsFromStateResponse
}
