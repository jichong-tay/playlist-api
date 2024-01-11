package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// newErrorResponse creates a new error response.
func newErrorResponse(statusCode int) gin.H {
	return gin.H{
		"message":     http.StatusText(statusCode),
		"status_code": statusCode,
	}
}
