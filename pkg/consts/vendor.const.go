package consts

import (
	"fmt"
	"os"

	vstructs "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendorpkg/structs"
)

type VendorNamesE string

const (
	Consensys VendorNamesE = "Consensys"
)

var VendorConfigMap = map[VendorNamesE]vstructs.Vendor{}

func InitializeVendorConfig() {
	VendorConfigMap[Consensys] = vstructs.Vendor{
		BaseURL: fmt.Sprintf("%s:%s", os.Getenv(string(CONSENSYS_CLIENT_HOST)), os.Getenv(string(CONSENSYS_CLIENT_PORT))),
	}
}
