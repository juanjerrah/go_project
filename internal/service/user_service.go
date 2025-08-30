package service

import (
	"github.com/juanjerrah/go-project/internal/models"
	"github.com/juanjerrah/go-project/internal/repository"
)

type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(id uint, req *models.UpdateUserRequest) (*models.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser implements UserService.
func (u *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := u.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete implements UserService.
func (u *userService) Delete(id uint) error {
	if err := u.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

// GetAll implements UserService.
func (u *userService) GetAll() ([]models.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID implements UserService.
func (u *userService) GetByID(id uint) (*models.User, error) {
	user, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements UserService.
func (u *userService) Update(id uint, req *models.UpdateUserRequest) (*models.User, error) {
	user, err := u.repo.GetByID(id)

	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	if err := u.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
