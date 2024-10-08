// Package middleware contains custom middleware for the application.
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// DurationMiddleware measures the duration of each request.
func DurationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Record the start time

		// Process the request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log the duration to the console
		log.Printf("Request processed in %s\n", duration)
	}
}
