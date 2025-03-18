package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs request method, path, status, and duration
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next() // Process request

		// Log details
		log.Printf("[%s] %s | %d | %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(startTime),
		)
	}
}
