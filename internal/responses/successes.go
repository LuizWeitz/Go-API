package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/internal/models"
)

func Ok[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, models.SuccessData[T]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	})
}

func Create[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusCreated, models.SuccessData[T]{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   data,
	})
}

func Update[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, models.SuccessData[T]{
		Code:   http.StatusOK,
		Status: "UPDATED",
		Data:   data,
	})
}

func Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Success{
		Code:   http.StatusOK,
		Status: "DELETED",
	})
}

func List[T any](ctx *gin.Context, data []T) {
	ctx.JSON(http.StatusOK, models.SuccessList[T]{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	})
}
