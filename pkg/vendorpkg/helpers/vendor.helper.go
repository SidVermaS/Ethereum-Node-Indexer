package helpers

// Checks whether the http response was successful or not
func CheckSuccess(statusCode int) bool {
	return statusCode > 199 && statusCode < 300
}