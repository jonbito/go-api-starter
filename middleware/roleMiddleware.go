package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// RoleMiddleware checks if the user has a role assigned, if not, the request is aborted
func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if claims["id"] == nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "You are not authorized to access this resource.",
			})
		}
		c.Next()
	}
}
