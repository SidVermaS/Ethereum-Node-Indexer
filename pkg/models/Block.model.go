package models

import "time"

type  Block struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Hash      string    `json:"hash" gorm:"unique;not null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
}
