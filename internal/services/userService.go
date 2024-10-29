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

type UserServiceImplementation struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImplementation{userRepository: userRepository}
}

func (usi *UserServiceImplementation) Create(user *models.User) (*models.User, error) {
	return usi.userRepository.Create(user)
}

func (usi *UserServiceImplementation) Delete(id uuid.UUID) error {
	return usi.userRepository.Delete(id)
}

func (usi *UserServiceImplementation) GetAll() ([]*models.User, error) {
	return usi.userRepository.GetAll()
}

func (usi *UserServiceImplementation) GetByID(id uuid.UUID) (*models.User, error) {
	return usi.userRepository.GetByID(id)
}

func (usi *UserServiceImplementation) Update(user *models.User) error {
	return usi.userRepository.Update(user)
}
