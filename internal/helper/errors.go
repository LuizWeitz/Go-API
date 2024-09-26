package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizweitz/go-api/internal/models"
)

func ErrDuplicatedKey(ctx *gin.Context, field string) {
	ctx.JSON(http.StatusConflict, models.Error{Code: http.StatusConflict, Message: field + " already exists", Status: "ALREADY_EXISTS"})
}

func ErrBindJson(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Message: err.Error(), Status: "BIND_JSON"})
}

func ErrNotFound(ctx *gin.Context, model string) {
	ctx.JSON(http.StatusNotFound, models.Error{Code: http.StatusNotFound, Message: model + " not found", Status: "NOT_FOUND"})
}

func ErrInternalServer(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, models.Error{Code: http.StatusInternalServerError, Message: "internal server error", Status: "INTERNAL"})
}

func ErrBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, models.Error{Code: http.StatusBadRequest, Message: message, Status: "BAD_REQUEST"})
}
