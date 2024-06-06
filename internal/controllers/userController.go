package controllers

import (
	"errors"
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

func (u *UserControllerImpl) Create(ctx *gin.Context) {

	var newUser *models.User

	if errBind := ctx.ShouldBindJSON(&newUser); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	user, errCreate := u.userService.Create(newUser)

	if errCreate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errCreate.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

func (u *UserControllerImpl) Delete(ctx *gin.Context) {

	userId, _ := uuid.Parse(ctx.Param("id"))

	_, errGet := u.userService.GetByID(userId)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		return
	}

	errDelete := u.userService.Delete(userId)

	if errDelete != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errDelete.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})

}

func (u *UserControllerImpl) GetAll(ctx *gin.Context) {

	var users []*models.User

	users, errGet := u.userService.GetAll()

	if errGet != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errGet.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)

}

func (u *UserControllerImpl) GetById(ctx *gin.Context) {

	userId, _ := uuid.Parse(ctx.Param("id"))

	var user *models.User

	user, errGet := u.userService.GetByID(userId)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		return
	}

	if errGet != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errGet.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (u *UserControllerImpl) Update(ctx *gin.Context) {

	userId, _ := uuid.Parse(ctx.Param("id"))

	_, errGet := u.userService.GetByID(userId)

	if errors.Is(errGet, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": errGet.Error()})
		return
	}

	var user *models.User

	if errBind := ctx.ShouldBindJSON(&user); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errBind.Error()})
		return
	}

	errUpdate := u.userService.Update(user)

	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errUpdate.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "success"})
}
