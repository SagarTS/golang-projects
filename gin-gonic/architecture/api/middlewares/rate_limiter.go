package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

func RateLimiter() gin.HandlerFunc{
	// create a store
	store := memory.NewStore()

	rate := limiter.Rate{
		Period: 3 * time.Second,
		Limit: 4,
	}

	limiter := limiter.New(store,rate)

	return func(c *gin.Context){
		// Check for rate of limit
		ctx, _ := limiter.Get(c,c.FullPath())

		if(ctx.Reached){
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many request, please try again sometime after",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}