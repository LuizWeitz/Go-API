package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/luizweitz/go-api/internal/helper"
	"github.com/luizweitz/go-api/internal/models"
	"github.com/luizweitz/go-api/internal/services"
	"gorm.io/gorm"
)

type UserController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserController(UserService services.UserService) UserController {
	return &UserControllerImpl{userService: UserService}
}

// @Summary		Search All Users
// @Router			/users [get]
// @Description	Search All Users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.User
// @Failure		400	{object}	models.Error	"Error"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImpl) GetAll(ctx *gin.Context) {

	var users []*models.User

	users, errGet := uci.userService.GetAll()

	if errGet != nil {
		helper.ErrInternalServer(ctx)
		return
	}

	ctx.JSON(http.StatusOK, users)

}

// @Summary		Search User By ID
// @Router			/users/{id} [get]
// @Description	Get User By ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	models.User
// @Failure		400	{object}	models.Error	"Error"
// @Failure		404	{object}	models.Error	"Eecord Not Found"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImpl) GetById(ctx *gin.Context) {

	userID, _ := uuid.Parse(ctx.Param("id"))

	var user *models.User

	user, errGet := uci.userService.GetByID(userID)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		helper.ErrNotFound(ctx, "user")
		return
	}

	if errGet != nil {
		helper.ErrInternalServer(ctx)
		return
	}

	ctx.JSON(http.StatusOK, user)

}

// @Summary		Create User
// @Router			/users [post]
// @Description	Create User With The Given Input Data
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			Input	body		models.User		true	"Create user object"
// @Success		200	{object}	models.User
// @Failure		400	{object}	models.Error	"Error"
// @Failure		409 {object}	models.Error	"Error Conflict"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImpl) Create(ctx *gin.Context) {

	var newUser *models.User

	if errBind := ctx.ShouldBindJSON(&newUser); errBind != nil {
		helper.ErrBindJson(ctx, errBind)
		return
	}

	user, errCreate := uci.userService.Create(newUser)

	if errCreate != nil {

		var pgErr *pgconn.PgError

		if errors.As(errCreate, &pgErr) {

			if pgErr.Code == "23505" {
				helper.ErrDuplicatedKey(ctx, pgErr.ConstraintName)
				return
			}

			helper.ErrBadRequest(ctx, pgErr.Message)
			return
		}

		helper.ErrInternalServer(ctx)
		return

	}

	ctx.JSON(http.StatusCreated, user)
}

// @Summary		Update User By ID
// @Router			/users [put]
// @Description	Update Data User By ID With The Given Input Data
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			Input	body		models.User		true	"Update user object"
// @Success		200		{string}	string			"Success"
// @Failure		400	{object}	models.Error	"Error"
// @Failure		404		{object}	models.Error	"Record Not Found"
// @Failure		409 {object}	models.Error	"Error Conflict"
// @Failure		500		{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImpl) Update(ctx *gin.Context) {

	var user *models.User

	if errBind := ctx.ShouldBindJSON(&user); errBind != nil {
		helper.ErrBindJson(ctx, errBind)
		return
	}

	_, errGet := uci.userService.GetByID(user.ID)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		helper.ErrNotFound(ctx, "user")
		return
	}

	errUpdate := uci.userService.Update(user)

	if errUpdate != nil {

		var pgErr *pgconn.PgError

		if errors.As(errUpdate, &pgErr) {

			if pgErr.Code == "23505" {
				helper.ErrDuplicatedKey(ctx, pgErr.ConstraintName)
				return
			}

			helper.ErrBadRequest(ctx, pgErr.Message)
			return
		}

		helper.ErrInternalServer(ctx)
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})
}

// @Summary		Delete User By ID
// @Router			/users/{id} [delete]
// @Description	Delete User By ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{string}		string      "Success"
// @Failure		400	{object}	models.Error	"Error"
// @Failure		404	{object}	models.Error	"Record Not Found"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImpl) Delete(ctx *gin.Context) {

	userID, _ := uuid.Parse(ctx.Param("id"))

	_, errGet := uci.userService.GetByID(userID)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		helper.ErrNotFound(ctx, "user")
		return
	}

	errDelete := uci.userService.Delete(userID)

	if errDelete != nil {
		helper.ErrInternalServer(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})

}
