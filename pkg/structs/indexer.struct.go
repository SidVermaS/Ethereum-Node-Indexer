package structs

type CalculateParticipatiRateStruct struct {
	MissedAttestations int
	Epochs             int
	SlotsPerEpoch      int
	ValidatorSetSize   int
}
