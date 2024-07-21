package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary		Create User
// @Router			/users [post]
// @Description	Create User With The Given Input Data
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			Input	body		models.User		true	"Create user object"
// @Success		200	{object}	models.User
// @Failure		400	{object}	models.Error	"error"
// @Failure		500	{object}	models.Error	"internal server error"
func (u *UserControllerImpl) Create(ctx *gin.Context) {

	var newUser *models.User

	if errBind := ctx.ShouldBindJSON(&newUser); errBind != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{Error: errBind.Error()})
		return
	}

	user, errCreate := u.userService.Create(newUser)

	if errCreate != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Error: errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

// @Summary		Delete User By ID
// @Router			/users/{id} [delete]
// @Description	Delete User By ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{string}		string      "success"
// @Failure		400	{object}	models.Error	"error"
// @Failure		404	{object}	models.Error	"record not found"
// @Failure		500	{object}	models.Error	"internal server error"
func (u *UserControllerImpl) Delete(ctx *gin.Context) {

	userId, _ := uuid.Parse(ctx.Param("id"))

	_, errGet := u.userService.GetByID(userId)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, models.Error{Error: "user not found"})
		return
	}

	errDelete := u.userService.Delete(userId)

	if errDelete != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Error: errDelete.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})

}

// @Summary		Search All Users Active
// @Router			/users [get]
// @Description	Search All Users Active
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.User
// @Failure		400	{object}	models.Error	"error"
// @Failure		500	{object}	models.Error	"internal server error"
func (u *UserControllerImpl) GetAll(ctx *gin.Context) {

	var users []*models.User

	users, errGet := u.userService.GetAll()

	if errGet != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Error: "internal server error"})
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
// @Failure		400	{object}	models.Error	"error"
// @Failure		404	{object}	models.Error	"record not found"
// @Failure		500	{object}	models.Error	"internal server error"
func (u *UserControllerImpl) GetById(ctx *gin.Context) {

	userId, _ := uuid.Parse(ctx.Param("id"))

	var user *models.User

	user, errGet := u.userService.GetByID(userId)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, models.Error{Error: "user not found"})
		return
	}

	if errGet != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Error: "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

// @Summary		Update User By ID
// @Router			/users/{id} [put]
// @Description	Update Data User By ID With The Given Input Data
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		string			true	"User ID"
// @Param			Input	body		models.User		true	"Create user object"
// @Success		200		{string}	string			"success"
// @Failure		400	{object}	models.Error	"error"
// @Failure		404		{object}	models.Error	"record not found"
// @Failure		500		{object}	models.Error	"internal server error"
func (u *UserControllerImpl) Update(ctx *gin.Context) {

	userId, _ := uuid.Parse(ctx.Param("id"))

	fmt.Println(userId)

	_, errGet := u.userService.GetByID(userId)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {

		ctx.JSON(http.StatusNotFound, models.Error{Error: errGet.Error()})
		return
	}

	var user *models.User

	if errBind := ctx.ShouldBindJSON(&user); errBind != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{Error: errBind.Error()})
		return
	}

	errUpdate := u.userService.Update(user)

	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{Error: errUpdate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})
}
