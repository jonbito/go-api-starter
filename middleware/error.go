package middleware

import (
	"github.com/gin-gonic/gin"
)

// ErrorMiddleware handles any errors that might happen
func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors)
	}
}
