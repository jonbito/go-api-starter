package middleware

import (
	"fmt"
	"go-api-starter/models"
	"go-api-starter/repository"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateAuthMiddleware creates authentication middleware for gin
func CreateAuthMiddleware(repo repository.IRepository, realm string, secret string, timeout time.Duration, maxRefresh time.Duration) *jwt.GinJWTMiddleware {
	// Setup Auth Middleware
	return &jwt.GinJWTMiddleware{
		Realm:      realm,
		Key:        []byte(secret),
		Timeout:    timeout,
		MaxRefresh: maxRefresh,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			var user models.User
			if db.Where("email = ?", email).First(&user).RecordNotFound() {
				return "", false
			}
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return "", false
			}

			return fmt.Sprint(user.ID), true
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
