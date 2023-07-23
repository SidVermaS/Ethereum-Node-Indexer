package models

type Validator struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	PublicKey string `json:"public_key" gorm:"uniqueIndex:idx_publickey;not null;type:varchar(200)"`
	Index uint64 `json:"index" gorm:"not null"`
}
