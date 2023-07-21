package consensysconsts

type ConsensysTopicsE string

const (
	Head                   ConsensysTopicsE = "head"
	Finalized_checkpoint   ConsensysTopicsE = "finalized_checkpoint"
	Chain_reorg            ConsensysTopicsE = "chain_reorg"
	Block                  ConsensysTopicsE = "block"
	Attestation            ConsensysTopicsE = "attestation"
	Voluntary_exit         ConsensysTopicsE = "voluntary_exit"
	Contribution_and_proof ConsensysTopicsE = "contribution_and_proof"
)
var AllConsensysTopics =[]ConsensysTopicsE{
	Head,
	Finalized_checkpoint,
	Chain_reorg,
	Block,
	Attestation,
	Voluntary_exit,
	Contribution_and_proof,
}