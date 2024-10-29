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

type UserRepositoryImplementation struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &UserRepositoryImplementation{db: db}
}

func (uri *UserRepositoryImplementation) Create(user *models.User) (*models.User, error) {

	if err := uri.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uri *UserRepositoryImplementation) Delete(id uuid.UUID) error {
	if err := uri.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (uri *UserRepositoryImplementation) GetAll() (users []*models.User, err error) {
	if err := uri.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (uri *UserRepositoryImplementation) GetByID(id uuid.UUID) (user *models.User, err error) {
	if err := uri.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (uri *UserRepositoryImplementation) Update(user *models.User) error {

	if err := uri.db.Model(&user).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
