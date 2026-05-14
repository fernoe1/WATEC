package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	windowSize = 10 * time.Second
	limit      = 5
)

func RateLimit(r *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "rl:" + c.ClientIP()
		count, _ := r.Get(c, key).Int()

		allowed := count < limit

		if !allowed {
			c.Status(http.StatusTooManyRequests)
			c.Abort()

			return
		}

		p := r.TxPipeline()
		p.Incr(c, key)
		p.ExpireNX(c, key, windowSize)

		c.Next()
	}
}
