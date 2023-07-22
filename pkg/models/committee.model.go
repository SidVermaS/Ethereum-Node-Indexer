package models

type Committee struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Eid uint `json:"eid" gorm:"not null"`
	Epoch Epoch `gorm:"foreignKey:Eid;not null"`
	Sid uint `json:"sid" gorm:"not null"`
	Slot Slot `gorm:"foreignKey:Sid;not null"`
}
