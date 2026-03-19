package httphelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONError sends a standard error response.
func JSONError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"error": err.Error()})
}

// JSONAccessToken sends only the access token (no refresh token).
func JSONAccessToken(c *gin.Context, accessToken string) {
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}
