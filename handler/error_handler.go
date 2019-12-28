package handler

import "error-handling-go/app_error"

func HandleErr(err error) string {
	if app_error.IsBadRequest(err) {
		return "Bad Request Handling"
	}
	if app_error.IsUnauthorized(err) {
		return "Unauthorized Handling"
	}
	if app_error.IsNotFound(err) {
		return "NotFound Handling"
	}
	return "Internal server error Handling"
}
