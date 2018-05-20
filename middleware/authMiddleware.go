package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// CreateAuthMiddleware creates authentication middleware for gin
func CreateAuthMiddleware(db *gorm.DB) *jwt.GinJWTMiddleware {
	// Setup Auth Middleware
	return &jwt.GinJWTMiddleware{
		Realm:      "my realm",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return email, false
			}

			if db.Table("users").Where("email = ?", email).Where("password = ?", hash).RecordNotFound() {
				return email, false
			}

			return email, true
		},
		Authorizator: func(email string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}
