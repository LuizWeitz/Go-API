package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/internal/models"
)

func Ok[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, models.SuccessData[T]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func Create[T any](c *gin.Context, data T) {
	c.JSON(http.StatusCreated, models.SuccessData[T]{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   data,
	})
}

func Update[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, models.SuccessData[T]{
		Code:   http.StatusOK,
		Status: "UPDATED",
		Data:   data,
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, models.Success{
		Code:   http.StatusOK,
		Status: "DELETED",
	})
}

func List[T any](c *gin.Context, data []T) {
	c.JSON(http.StatusOK, models.SuccessList[T]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	})
}
