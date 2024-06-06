package repositories

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/luizweitz/go-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uuid.UUID) (*models.User, error)
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

func (u *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {

	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Delete(id uuid.UUID) error {
	if err := u.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) GetAll() (users []*models.User, err error) {
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) GetById(id uuid.UUID) (user *models.User, err error) {
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Update(user *models.User) error {

	fmt.Print(time.Now())

	if errUpdate := u.db.Save(&models.User{ID: user.ID, Name: user.Name, Email: user.Email, Age: user.Age, City: user.City, UpdatedAt: time.Now(), CreatedAt: user.CreatedAt, DeletedAt: user.DeletedAt}).Error; errUpdate != nil {
		return errUpdate
	}

	return nil
}
