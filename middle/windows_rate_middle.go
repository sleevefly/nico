package middle

import (
	"github.com/gin-gonic/gin"
	"nico/limiter"
)

func SlidingWindowRateLimiterMiddleware(rateLimiter *limiter.GWindowsRate) gin.HandlerFunc {
	return func(c *gin.Context) {
		if rateLimiter.Allow(c.Request.RequestURI, true) {
			c.Next()
		} else {
			c.JSON(429, gin.H{"error": "请求太频繁，请稍后重试"})
			c.Abort()
		}
	}
}
