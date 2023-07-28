package models

type Epoch struct {
	ID       uint  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Period   uint  `json:"period" gorm:"uniqueIndex:idx_period;not null"`
	BlockId *uint  `json:"block_id,omitempty"`
	Block    *Block `json:"block,omitempty" gorm:"foreignKey:block_id;not null"`
}

// It's an interface needed for overriding the table name by GoORM
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Epoch to `epochs`
func (Epoch) TableName() string {
	return "epochs"
}
