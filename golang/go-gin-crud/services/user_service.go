// Package services contains the business logic for user-related operations.
package services

import (
	"github.com/gitish/polyglot_training/repositories"

	"github.com/gitish/polyglot_training/models"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	CreateUser(user models.User) models.User
	GetAllUsers() []models.User
	GetUserByID(id int) (models.User, bool)
	UpdateUser(user models.User) (models.User, bool)
	DeleteUser(id int) bool
}

// userServiceImpl implements the UserService interface.
type userServiceImpl struct {
	repo repositories.UserRepository
}

// NewUserService creates a new instance of userServiceImpl.
func NewUserService(repo repositories.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

// CreateUser creates a new user.
func (s *userServiceImpl) CreateUser(user models.User) models.User {
	return s.repo.CreateUser(user)
}

// GetAllUsers retrieves all users.
func (s *userServiceImpl) GetAllUsers() []models.User {
	return s.repo.GetAllUsers()
}

// GetUserByID retrieves a user by their ID.
func (s *userServiceImpl) GetUserByID(id int) (models.User, bool) {
	return s.repo.GetUserByID(id)
}

// UpdateUser updates an existing user.
func (s *userServiceImpl) UpdateUser(user models.User) (models.User, bool) {
	return s.repo.UpdateUser(user)
}

// DeleteUser deletes a user by their ID.
func (s *userServiceImpl) DeleteUser(id int) bool {
	return s.repo.DeleteUser(id)
}
