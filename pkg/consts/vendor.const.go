package consts

import (
	"fmt"
	"os"

	vstructs "github.com/SidVermaS/Ethereum-Consensus/pkg/vendorpkg/structs"
)

type VendorNamesE string

const (
	Consensys VendorNamesE = "Consensys"
)

var VendorConfigMap = map[VendorNamesE]vstructs.Vendor{Consensys: {
	BaseURL: os.Getenv(fmt.Sprintf("%s%s", string(CONSENSYS_CLIENT_HOST), string(CONSENSYS_CLIENT_PORT))),
}}

