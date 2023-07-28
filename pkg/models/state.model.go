package models

type State struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	StateStored string `json:"state_stored" gorm:"uniqueIndex:idx_statestored;not null;type:varchar(200)"`
	BlockId     *uint   `json:"block_id,omitempty"`
	Block       *Block  `json:"block,omitempty" gorm:"foreignKey:block_id;not null"`
	EpochId     *uint   `json:"epoch_id,omitempty" gorm:"not null"`
	Epoch       *Epoch  `json:"epoch,omitempty" gorm:"foreignKey:EpochId;not null"`
}
