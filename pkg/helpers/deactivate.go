package helpers

// Closing all connections, event listeners, etc
func Deactivate() {
	CloseDBConnection()
}

// The gorm DB Connection is closed
func CloseDBConnection() {
	db, _ := Repository.DB.DB()
	db.Close()
}
