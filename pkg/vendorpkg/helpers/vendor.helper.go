package helpers

func CheckSuccess(statusCode int) bool {
	return statusCode > 199 && statusCode < 300
}