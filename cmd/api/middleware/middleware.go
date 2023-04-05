package middleware

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

//
func MiddlewareVerificationToken() gin.HandlerFunc {
	//add token in the header
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN"){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}



