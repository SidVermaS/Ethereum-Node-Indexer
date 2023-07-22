package models

type Validator struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	PublicKey string `json:"public_key" gorm:"type:varchar(200);unique_index;not null"`
	IsSlashed bool `json:"is_slashed" gorm:"not null"`
	Status string `json:"status"`
}
