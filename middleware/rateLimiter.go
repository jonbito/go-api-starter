package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api-starter/middleware/limiter"
	mgin "go-api-starter/middleware/limiter/drivers/middleware/gin"
	"go-api-starter/middleware/limiter/drivers/store/memory"
	"time"
)

// RateLimiter adds a rate limit to GIN requests
func RateLimiter(period time.Duration, limit int64) gin.HandlerFunc {
	rate := limiter.Rate{
		Period: period,
		Limit:  limit,
	}

	store := memory.NewStore()

	middleware := mgin.NewMiddleware(limiter.New(store, rate))
	return middleware
}
