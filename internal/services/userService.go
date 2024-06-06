package services

import (
	"github.com/google/uuid"
	"github.com/luizweitz/go-api/internal/models"
	"github.com/luizweitz/go-api/internal/repositories"
)

type UserService interface {
	GetAll() ([]*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
}

type UserServiceImpl struct {
	userrepositories repositories.UserRepository
}

func NewUserService(Userrepositories repositories.UserRepository) UserService {
	return &UserServiceImpl{userrepositories: Userrepositories}
}

func (us *UserServiceImpl) Create(user *models.User) (*models.User, error) {
	return us.userrepositories.Create(user)
}

func (us *UserServiceImpl) Delete(id uuid.UUID) error {
	return us.userrepositories.Delete(id)
}

func (us *UserServiceImpl) GetAll() ([]*models.User, error) {
	return us.userrepositories.GetAll()
}

func (us *UserServiceImpl) GetByID(id uuid.UUID) (*models.User, error) {
	return us.userrepositories.GetById(id)
}

func (us *UserServiceImpl) Update(user *models.User) error {
	return us.userrepositories.Update(user)
}
