package models

type Epoch struct {
	ID    uint  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Period  uint  `json:"period" gorm:"uniqueIndex:idx_period;not null"`
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
