package models

type Slot struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Index uint64 `json:"index" gorm:"not null"`
	Eid uint `json:"eid" gorm:"not null"`
	Epoch Epoch `gorm:"foreignKey:Eid;not null"`
	StateId   uint      `json:"state_id" gorm:"not null"`
	State     State     `gorm:"foreignKey:StateId;not null"`
}
