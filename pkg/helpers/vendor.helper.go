package helpers

import (
	"github.com/SidVermaS/Ethereum-Consensus/pkg/consts"
	vendors "github.com/SidVermaS/Ethereum-Consensus/pkg/vendors/consensys"
)

func GetVendor(VendorName consts.VendorNamesE) *vendors.Consensys {
	var ConsensysVendor *vendors.Consensys = &vendors.Consensys{}
	ConsensysVendor.Vendor = consts.VendorConfigMap[VendorName]
	return ConsensysVendor
}
