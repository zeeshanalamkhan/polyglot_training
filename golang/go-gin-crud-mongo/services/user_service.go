package services

import (
	"fmt"

	"github.com/gitish/polyglot_training/models"
	"github.com/gitish/polyglot_training/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id primitive.ObjectID) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id primitive.ObjectID) error
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
func (s *userServiceImpl) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

// GetAllUsers retrieves all users.
func (s *userServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers(), nil
}

// GetUserByID retrieves a user by their ID.
func (s *userServiceImpl) GetUserByID(id primitive.ObjectID) (models.User, error) {
	user, exists := s.repo.GetUserByID(id)
	if !exists {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

// UpdateUser updates an existing user.
func (s *userServiceImpl) UpdateUser(user models.User) (models.User, error) {
	updatedUser, exists := s.repo.UpdateUser(user)
	if !exists {
		return models.User{}, fmt.Errorf("user not found")
	}
	return updatedUser, nil
}

// DeleteUser deletes a user by their ID.
func (s *userServiceImpl) DeleteUser(id primitive.ObjectID) error {
	if !s.repo.DeleteUser(id) {
		return fmt.Errorf("user not found")
	}
	return nil
}
