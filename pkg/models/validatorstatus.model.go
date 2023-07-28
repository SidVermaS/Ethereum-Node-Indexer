package models

type ValidatorStatus struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	IsSlashed   bool      `json:"is_slashed" gorm:"not null"`
	Status      string    `json:"status"`
	ValidatorId *uint      `json:"validator_id,omitempty" gorm:"not null"`
	Validator   *Validator `json:"validator,omitempty" gorm:"foreignKey:ValidatorId;not null"`
	StateId     *uint      `json:"state_id,omitempty" gorm:"not null"`
	State       *State     `json:"state,omitempty" gorm:"foreignKey:StateId;not null"`
	EpochId     *uint      `json:"epoch_id,omitempty" gorm:"not null"`
	Epoch       *Epoch     `json:"epoch,omitempty" gorm:"foreignKey:EpochId;not null"`
	BlockId     *uint      `json:"block_id,omitempty"`
	Block       *Block     `json:"block,omitempty" gorm:"foreignKey:block_id;not null"`
}
