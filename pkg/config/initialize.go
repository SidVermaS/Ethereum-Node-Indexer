package config

import (
	"os"

	"github.com/SidVermaS/Ethereum-Consensus-Layer/pkg/consts"
	"github.com/SidVermaS/Ethereum-Consensus-Layer/pkg/migrations"
	"github.com/SidVermaS/Ethereum-Consensus-Layer/pkg/types/structs"
	"github.com/joho/godotenv"
)

func InitializeDB() {
	dbConfig := &structs.DbConfig{
		Host:     os.Getenv(string(consts.POSTGRES_HOST)),
		User:     os.Getenv(string(consts.POSTGRES_USER)),
		Password: os.Getenv(string(consts.POSTGRES_PASSWORD)),
		DBName:   os.Getenv(string(consts.POSTGRES_DB)),
		Port:     os.Getenv(string(consts.POSTGRES_PORT)),
		SSLMode:  os.Getenv(string(consts.POSTGRES_SSL_MODE)),
	}
	CreateConnection(dbConfig)
}

func InitializeAll() {
	// Load the .env file
	godotenv.Load(".env")

	// Initializing the Database
	InitializeDB()
	// It needs to be executed only for the first time
	migrations.InitialMigration(Repository.DB)
}
