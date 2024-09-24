package helpers

import (
	"errors"
)

var (
	ErrEmailAlreadyRegistered    = errors.New("email has already been registered")
	ErrUserNameAlreadyRegistered = errors.New("username has already been registered")
	ErrAgeMustBeMore             = errors.New("age must be more than 8")
	ErrAccountNotFound           = errors.New("account not found")
)

// // Helper function to handle errors
// func RespondWithError(c *gin.Context, statusCode int, errorTitle string, errorMsg string) {
// 	c.AbortWithStatusJSON(statusCode, gin.H{
// 		"error":   errorTitle,
// 		"message": errorMsg,
// 	})
// }

// // Helper function to handle JSON binding errors
// func RespondBindJSONError(c *gin.Context, errorMsg string) {
// 	RespondWithError(c, http.StatusBadRequest, "Invalid request payload", errorMsg)
// }

// // Helper function to handle validation errors
// func RespondValidatorError(c *gin.Context, errorMsg string) {
// 	RespondWithError(c, http.StatusBadRequest, "Validation failed", errorMsg)
// }

// // Helper fucntion to hendle ID invalid errors
// func RespondIDInvalidError(c *gin.Context, errorMsg string) {
// 	RespondWithError(c, http.StatusBadRequest, "ID invalid", errorMsg)
// }
