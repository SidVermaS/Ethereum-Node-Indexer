package configs

// Closing all connections, event listeners, etc
func Deactivate() {
	CloseDBConnection()
}
