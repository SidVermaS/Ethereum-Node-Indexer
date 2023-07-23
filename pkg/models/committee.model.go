package models

type Committee struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Eid uint `json:"eid" gorm:"not null"`
	Epoch Epoch `gorm:"foreignKey:Eid;not null"`
	SlotId uint `json:"slot_id" gorm:"not null"`
	Slot Slot `gorm:"foreignKey:SlotId;not null"`
	Vid uint `json:"vid" gorm:"not null"`
	Validator Validator `gorm:"foreignKey:Vid;not null"`
}
