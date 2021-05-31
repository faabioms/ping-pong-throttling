package router

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	
	limiter "github.com/faabioms/ping-pong-throttling/gin-limiter"
)

func Router() *gin.Engine {
	r := gin.Default()
	//{ "message": "request throttled request", "throttle_age": int }
	lmt := limiter.NewRateLimiter(time.Minute, 10, func(c *gin.Context) (string, error) {
		key := c.Request.Header.Get("x-secret-key")
		if key != "" {
			return key, nil
		}
		return "", errors.New("key is missing")
	})

	lmt2 := limiter.NewRateLimiter(time.Second, 2, func(c *gin.Context) (string, error) {
		key := c.Request.Header.Get("x-secret-key")
		if key != "" {
			return key, nil
		}
		return "", errors.New("key is missing")
	})

	r.POST("ping", lmt.Middleware(), lmt2.Middleware(), func(c *gin.Context) {
		// parse json
		var request struct {
			Value string `json:"request"`
		}
		
		type header struct {
			AccessKey string `header:"x-secret-key"`
		}

		h := header{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		if c.Bind(&request) == nil {
			if request.Value == "ping" || h.AccessKey != "" {
				c.JSON(http.StatusOK, gin.H{"response": "pong"})
			}
		}
	})
	return r
}