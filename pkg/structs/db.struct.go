package structs

import "gorm.io/gorm"

// struct for the configuration of the database 
type DbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type DBRepository struct {
	DB *gorm.DB
}