package models

type Slot struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Index   uint64 `json:"index" gorm:"not null"`
	EpochId *uint   `json:"epoch_id,omitempty" gorm:"not null"`
	Epoch   *Epoch  `json:"epoch,omitempty" gorm:"foreignKey:EpochId;not null"`
	StateId *uint   `json:"state_id,omitempty" gorm:"not null"`
	State   *State  `json:"state,omitempty" gorm:"foreignKey:StateId;not null"`
	BlockId *uint   `json:"block_id,omitempty"`
	Block   *Block  `json:"block,omitempty" gorm:"foreignKey:block_id;not null"`
}
