package models

type ValidatorStatus struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	IsSlashed bool      `json:"is_slashed" gorm:"not null"`
	Status    string    `json:"status"`
	Vid       uint      `json:"vid" gorm:"not null"`
	Validator Validator `gorm:"foreignKey:Vid;not null"`
	StateId   uint      `json:"state_id" gorm:"not null"`
	State     State     `gorm:"foreignKey:StateId;not null"`
	Eid       uint      `json:"eid" gorm:"not null"`
	Epoch     Epoch     `gorm:"foreignKey:Eid;not null"`
}
