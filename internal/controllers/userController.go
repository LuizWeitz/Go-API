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
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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
func (uci *UserControllerImplementation) GetAll(c *gin.Context) {

	var users []*models.User

	users, errGet := uci.userService.GetAll()

	if errGet != nil {
		responses.ErrorInternalServer(c)
		return
	}

	responses.List(c, users)

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
func (uci *UserControllerImplementation) GetById(c *gin.Context) {

	userID, errParse := uuid.Parse(c.Param("id"))

	if errParse != nil {

		responses.ErrorBadRequest(c, "UUID invalid")
		return
	}

	var user *models.User

	user, errGet := uci.userService.GetByID(userID)

	if errGet != nil {

		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			responses.ErrorNotFound(c, "user")
			return
		}

		responses.ErrorInternalServer(c)
		return
	}

	if !c.IsAborted() {
		responses.Ok(c, user)
		return
	}

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
func (uci *UserControllerImplementation) Create(c *gin.Context) {

	var newUser *models.User

	if errBind := c.ShouldBindJSON(&newUser); errBind != nil {
		responses.ErrorBindJson(c, errBind)
		return
	}

	user, errCreate := uci.userService.Create(newUser)

	if errCreate != nil {

		var pgErr *pgconn.PgError

		if errors.As(errCreate, &pgErr) {

			if pgErr.Code == "23505" {
				responses.ErrorDuplicatedKey(c, pgErr.ConstraintName)
				return
			}

			responses.ErrorBadRequest(c, pgErr.Message)
			return
		}

		responses.ErrorInternalServer(c)
		return
	}

	responses.Create(c, user)
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
func (uci *UserControllerImplementation) Update(c *gin.Context) {

	var user *models.User

	if errBind := c.ShouldBindJSON(&user); errBind != nil {
		responses.ErrorBindJson(c, errBind)
		return
	}

	if _, errGet := uci.userService.GetByID(user.ID); errGet != nil {
		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			responses.ErrorNotFound(c, "user")
			return
		}

		responses.ErrorInternalServer(c)
		return
	}

	if errUpdate := uci.userService.Update(user); errUpdate != nil {

		var pgErr *pgconn.PgError

		if errors.As(errUpdate, &pgErr) {

			if pgErr.Code == "23505" {
				responses.ErrorDuplicatedKey(c, pgErr.ConstraintName)
				return
			}

			responses.ErrorBadRequest(c, pgErr.Message)
			return
		}

		responses.ErrorInternalServer(c)
		return

	}

	responses.Update(c, user)
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
func (uci *UserControllerImplementation) Delete(c *gin.Context) {

	userID, errParse := uuid.Parse(c.Param("id"))

	if errParse != nil {
		responses.ErrorBadRequest(c, "UUID invalid")
		return
	}

	if _, errGet := uci.userService.GetByID(userID); errGet != nil {

		if errors.Is(errGet, gorm.ErrRecordNotFound) {
			responses.ErrorNotFound(c, "user")
			return
		}

		responses.ErrorInternalServer(c)
		return

	}

	if errDelete := uci.userService.Delete(userID); errDelete != nil {
		responses.ErrorInternalServer(c)
		return
	}

	responses.Delete(c)

}
