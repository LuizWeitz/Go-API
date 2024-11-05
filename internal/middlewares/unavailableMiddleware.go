package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/internal/models"
	"gorm.io/gorm"
)

type UnavailableMiddleware interface {
	CheckDatabase() gin.HandlerFunc
}

type UnavailableMiddlewareImpl struct {
	db *gorm.DB
}

func NewUnavailableMiddleware(db *gorm.DB) UnavailableMiddleware {
	return &UnavailableMiddlewareImpl{db: db}
}

func (umi *UnavailableMiddlewareImpl) CheckDatabase() gin.HandlerFunc {

	return func(c *gin.Context) {

		database, errConnect := umi.db.DB()

		if errConnect != nil {
			log.Println("Error Connect", errConnect)
			c.JSON(http.StatusServiceUnavailable, models.Error{
				Code:    http.StatusServiceUnavailable,
				Message: "service unavailable",
				Status:  "UNAVAILABLE",
			})
			c.Abort()
			return
		}

		errPing := database.Ping()

		if errPing != nil {
			log.Println("Error Ping :", errPing)
			c.JSON(http.StatusServiceUnavailable, models.Error{
				Code:    http.StatusServiceUnavailable,
				Message: "service unavailable",
				Status:  "UNAVAILABLE",
			})
			c.Abort()
			return
		}

		c.Next()

	}

}
