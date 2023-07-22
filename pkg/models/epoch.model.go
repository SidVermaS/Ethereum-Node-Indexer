package models

type Epoch struct {
	ID    uint  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Root  uint  `json:"root" gorm:"->;<-:create"`
	Bid   uint  `json:"bid"`
	Block Block `gorm:"foreignKey:Bid;not null"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Epoch to `epochs`
func (Epoch) TableName() string {
	return "epochs"
}
