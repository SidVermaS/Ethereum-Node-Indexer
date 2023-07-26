package consensysconsts

type ConsensysTopicsE string
type StateIdsE string

// These consts are the parameter values used in the Consensys APIs

const (
	Head                   ConsensysTopicsE = "head"
	Finalized_checkpoint   ConsensysTopicsE = "finalized_checkpoint"
	Chain_reorg            ConsensysTopicsE = "chain_reorg"
	Block                  ConsensysTopicsE = "block"
	Attestation            ConsensysTopicsE = "attestation"
	Voluntary_exit         ConsensysTopicsE = "voluntary_exit"
	Contribution_and_proof ConsensysTopicsE = "contribution_and_proof"
)

var AllConsensysTopics = []ConsensysTopicsE{
	Head,
	Finalized_checkpoint,
	Chain_reorg,
	Block,
	Attestation,
	Voluntary_exit,
	Contribution_and_proof,
}

const (
	Genesis   StateIdsE = "genesis"
	Finalized StateIdsE = "finalized"
	Justified StateIdsE = "justified"
)
type ValidatorStatusE string
const (
	ActiveOngoing ValidatorStatusE = "active_ongoing" 
	PendingInitialized ValidatorStatusE = "pending_initialized" 
)