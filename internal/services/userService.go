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

func (usi *UserServiceImpl) Create(user *models.User) (*models.User, error) {
	return usi.userRepository.Create(user)
}

func (usi *UserServiceImpl) Delete(id uuid.UUID) error {
	return usi.userRepository.Delete(id)
}

func (usi *UserServiceImpl) GetAll() ([]*models.User, error) {
	return usi.userRepository.GetAll()
}

func (usi *UserServiceImpl) GetByID(id uuid.UUID) (*models.User, error) {
	return usi.userRepository.GetByID(id)
}

func (usi *UserServiceImpl) Update(user *models.User) error {
	return usi.userRepository.Update(user)
}
