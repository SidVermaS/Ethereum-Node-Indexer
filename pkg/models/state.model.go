package models

type State struct {
	ID    uint  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	StateStored string `json:"state_stored" gorm:"uniqueIndex:idx_statestored;not null;type:varchar(200)"`
	Bid   uint  `json:"bid"`
	Block Block `gorm:"foreignKey:Bid;not null"`
	Eid uint `json:"eid" gorm:"not null"`
	Epoch Epoch `gorm:"foreignKey:Eid;not null"`
}
