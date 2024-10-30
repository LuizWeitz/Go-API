package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/internal/models"
)

func ErrorDuplicatedKey(c *gin.Context, field string) {
	c.JSON(http.StatusConflict, models.Error{
		Code:    http.StatusConflict,
		Message: field + " already exists",
		Status:  "ALREADY_EXISTS",
	})
}

func ErrorBindJson(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, models.Error{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
		Status:  "BIND_JSON",
	})
}

func ErrorNotFound(c *gin.Context, model string) {
	c.JSON(http.StatusNotFound, models.Error{
		Code:    http.StatusNotFound,
		Message: model + " not found",
		Status:  "NOT_FOUND",
	})
}

func ErrorInternalServer(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, models.Error{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Status:  "INTERNAL_SERVER_ERROR",
	})
}

func ErrorBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, models.Error{
		Code:    http.StatusBadRequest,
		Message: message,
		Status:  "BAD_REQUEST",
	})
}
