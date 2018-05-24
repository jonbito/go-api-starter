package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-api-starter/models"
)

// RoleMiddleware checks if the user has a role assigned, if not, the request is aborted
func RoleMiddleware(db *gorm.DB, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if claims["id"] == nil {
			errorOut(c)
			return
		}

		var user models.User
		if db.Find(&user, claims["id"]).Error != nil {
			errorOut(c)
			return
		}

		if user.Role != role {
			errorOut(c)
			return
		}

		c.Next()
	}
}

func errorOut(c *gin.Context) {
	c.AbortWithStatusJSON(403, gin.H{
		"error": "You are not authorized to access this resource.",
	})
}
