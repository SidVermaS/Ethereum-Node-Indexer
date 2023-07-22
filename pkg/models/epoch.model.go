package models

type Epoch struct {
	ID    uint  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Epoch  uint  `json:"epoch" gorm:"unique_index;not null"`
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
