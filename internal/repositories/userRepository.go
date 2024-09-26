package repositories

import (
	"github.com/google/uuid"
	"github.com/luizweitz/go-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uuid.UUID) (*models.User, error)
	GetAll() ([]*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &UserRepositoryImpl{db: db}
}

func (uci *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {

	if err := uci.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uci *UserRepositoryImpl) Delete(id uuid.UUID) error {
	if err := uci.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (uci *UserRepositoryImpl) GetAll() (users []*models.User, err error) {
	if err := uci.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (uci *UserRepositoryImpl) GetByID(id uuid.UUID) (user *models.User, err error) {
	if err := uci.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uci *UserRepositoryImpl) Update(user *models.User) error {

	if err := uci.db.Model(&user).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
