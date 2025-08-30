package repository

import (
	"github.com/juanjerrah/go-project/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create implements UserRepository.
func (u *userRepository) Create(user *models.User) error {
	return u.db.Create(user).Error
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id uint) error {
	return u.db.Delete(&models.User{}, id).Error
}

// GetAll implements UserRepository.
func (u *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	return users, err
}

// GetByID implements UserRepository.
func (u *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	return &user, err
}

// Update implements UserRepository.
func (u *userRepository) Update(user *models.User) error {
	return u.db.Save(user).Error
}

