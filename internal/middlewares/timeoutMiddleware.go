package middlewares

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/internal/models"
)

func Timeout() gin.HandlerFunc {

	timeoutStr := os.Getenv("TIMEOUT_SECONDS")
	timeoutValue, err := strconv.Atoi(timeoutStr)

	if timeoutValue <= 0 {
		log.Printf("Error : invalid timeout negative value, defaulting to 3 seconds")
		timeoutValue = 3
	}

	if err != nil {
		log.Printf("Error : invalid timeout value, defaulting to 3 seconds")
		timeoutValue = 3
	}

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeoutValue)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.JSON(http.StatusGatewayTimeout, models.Error{
				Code:    http.StatusGatewayTimeout,
				Message: "request timeout",
				Status:  "TIMEOUT",
			})
		}),
	)
}
