package models

type Block struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Root string `json:"root" gorm:"unique;not null"`
}
