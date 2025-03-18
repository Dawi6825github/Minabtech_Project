package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware handles panics and returns a 500 response
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
