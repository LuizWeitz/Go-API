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
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (u *UserServiceImpl) Create(user *models.User) (*models.User, error) {
	return u.userRepository.Create(user)
}

func (u *UserServiceImpl) Delete(id uuid.UUID) error {
	return u.userRepository.Delete(id)
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	return u.userRepository.GetAll()
}

func (u *UserServiceImpl) GetByID(id uuid.UUID) (*models.User, error) {
	return u.userRepository.GetByID(id)
}

func (u *UserServiceImpl) Update(user *models.User) error {
	return u.userRepository.Update(user)
}
