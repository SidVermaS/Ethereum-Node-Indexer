package structs

// Struct with attributes needed for the calculation of the paraticipation rate
type CalculateParticipatiRateStruct struct {
	MissedAttestations int
	Epochs             int
	SlotsPerEpoch      int
	ValidatorSetSize   int
}
