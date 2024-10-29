package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/luizweitz/go-api/internal/models"
	"github.com/luizweitz/go-api/internal/responses"
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

type UserControllerImplementation struct {
	userService services.UserService
}

func NewUserController(UserService services.UserService) UserController {
	return &UserControllerImplementation{userService: UserService}
}

// @Summary		Search All Users
// @Router			/users [get]
// @Description	Search All Users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200	{object}    models.SuccessList[models.User] "OK"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImplementation) GetAll(ctx *gin.Context) {

	var users []*models.User

	users, errGet := uci.userService.GetAll()

	if errGet != nil {
		responses.ErrorInternalServer(ctx)
		return
	}

	responses.List(ctx, users)

}

// @Summary		Search User By ID
// @Router			/users/{id} [get]
// @Description	Get User By ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	models.SuccessData[models.User] "OK"
// @Failure		400	{object}	models.Error	"Bad Request"
// @Failure		404	{object}	models.Error	"Not Found"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImplementation) GetById(ctx *gin.Context) {

	userID, errParse := uuid.Parse(ctx.Param("id"))

	if errParse != nil {
		responses.ErrorBadRequest(ctx, "UUID invalid")
		return
	}

	var user *models.User

	user, errGet := uci.userService.GetByID(userID)

	if errGet != nil {

		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			responses.ErrorNotFound(ctx, "user")
			return
		}

		responses.ErrorInternalServer(ctx)
		return
	}

	responses.Ok(ctx, user)

}

// @Summary		Create User
// @Router			/users [post]
// @Description	Create User With The Given Input Data
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			Input	body		models.User		true	"Create user object"
// @Success		201	{object}	models.SuccessData[models.User] "Created"
// @Failure		400	{object}	models.Error	"Bad Request"
// @Failure		409 {object}	models.Error	"Error Conflict"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImplementation) Create(ctx *gin.Context) {

	var newUser *models.User

	if errBind := ctx.ShouldBindJSON(&newUser); errBind != nil {
		responses.ErrorBindJson(ctx, errBind)
		return
	}

	user, errCreate := uci.userService.Create(newUser)

	if errCreate != nil {

		var pgErr *pgconn.PgError

		if errors.As(errCreate, &pgErr) {

			if pgErr.Code == "23505" {
				responses.ErrorDuplicatedKey(ctx, pgErr.ConstraintName)
				return
			}

			responses.ErrorBadRequest(ctx, pgErr.Message)
			return
		}

		responses.ErrorInternalServer(ctx)
		return
	}

	responses.Create(ctx, user)
}

// @Summary		Update User By ID
// @Router			/users [put]
// @Description	Update Data User By ID With The Given Input Data
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			Input	body		models.User		true	"Update user object"
// @Success		200	{object}	models.SuccessData[models.User] "Updated"
// @Failure		400	{object}	models.Error	"Bad Request"
// @Failure		404		{object}	models.Error	"Not Found"
// @Failure		409 {object}	models.Error	"Error Conflict"
// @Failure		500		{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImplementation) Update(ctx *gin.Context) {

	var user *models.User

	if errBind := ctx.ShouldBindJSON(&user); errBind != nil {
		responses.ErrorBindJson(ctx, errBind)
		return
	}

	if _, errGet := uci.userService.GetByID(user.ID); errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			responses.ErrorNotFound(ctx, "user")
			return
		}

		responses.ErrorInternalServer(ctx)
		return
	}

	if errUpdate := uci.userService.Update(user); errUpdate != nil {

		var pgErr *pgconn.PgError

		if errors.As(errUpdate, &pgErr) {

			if pgErr.Code == "23505" {
				responses.ErrorDuplicatedKey(ctx, pgErr.ConstraintName)
				return
			}

			responses.ErrorBadRequest(ctx, pgErr.Message)
			return
		}

		responses.ErrorInternalServer(ctx)
		return

	}

	responses.Update(ctx, user)
}

// @Summary		Delete User By ID
// @Router			/users/{id} [delete]
// @Description	Delete User By ID
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	models.Success[models.User] "Deleted"
// @Failure		400	{object}	models.Error	"Bed Request"
// @Failure		404	{object}	models.Error	"Not Found"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (uci *UserControllerImplementation) Delete(ctx *gin.Context) {

	userID, errParse := uuid.Parse(ctx.Param("id"))

	if errParse != nil {
		responses.ErrorBadRequest(ctx, "UUID invalid")
		return
	}

	if _, errGet := uci.userService.GetByID(userID); errGet != nil {

		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			responses.ErrorNotFound(ctx, "user")
			return
		}

		responses.ErrorInternalServer(ctx)
		return

	}

	if errDelete := uci.userService.Delete(userID); errDelete != nil {
		responses.ErrorInternalServer(ctx)
		return
	}

	responses.Delete(ctx)

}
