package consensys

import (
	"encoding/json"
	"fmt"
	"log"

	vstructs "github.com/SidVermaS/Ethereum-Consensus/pkg/vendorpkg/structs"

	"github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys/consensysstructs"
	consensysconsts "github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys/consts"
)

type Consensys struct {
	Vendor vstructs.Vendor
	Topics []string
}

func (consensys *Consensys) GetValidatorsFromState(stateId consensysconsts.StateIdsE) (*consensysstructs.GetValidatorsFromStateResponse) {
	statusCode, response, _ := consensys.Vendor.CallAPI(&vstructs.APIRequest{Url: fmt.Sprintf("/eth/v1/beacon/states/%s/validators", string(stateId))})
	fmt.Println("statusCode: ", statusCode)
	var getValidatorsFromStateResponse *consensysstructs.GetValidatorsFromStateResponse
	err := json.Unmarshal(response, &getValidatorsFromStateResponse)
	if err != nil {
		log.Printf("~~~ Error GetValidatorsFromState() %s\n", err)
	}
	return getValidatorsFromStateResponse
}
