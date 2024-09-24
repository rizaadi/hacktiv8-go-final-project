package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse represents the structure of the API response
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// RespondWithSuccess sends a success response
func RespondWithSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: message,
		Data:    data,
		Error:   nil,
	})
}

// RespondWithError sends an error response
func RespondWithError(c *gin.Context, status int, message string, err error) {
	c.AbortWithStatusJSON(status, APIResponse{
		Status:  http.StatusText(status),
		Message: message,
		Data:    nil,
		Error:   err.Error(),
	})
}

// Helper function to handle JSON binding errors
func RespondBindJSONError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: "Invalid request payload",
		Data:    nil,
		Error:   err.Error(),
	})

}

// Helper function to handle validation errors
func RespondValidatorError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: "Validation failed",
		Data:    nil,
		Error:   err.Error(),
	})
}

// Helper fucntion to hendle ID invalid errors
func RespondIDInvalidError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, APIResponse{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: "ID invalid",
		Data:    nil,
		Error:   err.Error(),
	})

}
