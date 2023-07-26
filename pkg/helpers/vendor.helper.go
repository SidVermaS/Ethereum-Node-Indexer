package helpers

import (
	"github.com/SidVermaS/Ethereum-Node-Indexer/pkg/consts"
	vendors "github.com/SidVermaS/Ethereum-Node-Indexer/pkg/vendors/consensys"
)
// Returns a vendor struct and initializes the configurations
func GetVendor(VendorName consts.VendorNamesE) *vendors.Consensys {
	var ConsensysVendor *vendors.Consensys = &vendors.Consensys{}
	ConsensysVendor.Vendor = consts.VendorConfigMap[VendorName]

	return ConsensysVendor
}
