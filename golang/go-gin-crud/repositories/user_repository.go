// Package repositories contains the repository layer for interacting with the in-memory database.
package repositories

import "github.com/gitish/polyglot_training/models"

// UserRepository defines the interface for user-related database operations.
type UserRepository interface {
	CreateUser(user models.User) models.User
	GetAllUsers() []models.User
	GetUserByID(id int) (models.User, bool)
	UpdateUser(user models.User) (models.User, bool)
	DeleteUser(id int) bool
}

// userRepositoryImpl implements the UserRepository interface using an in-memory database.
type userRepositoryImpl struct {
	users  map[int]models.User
	lastID int
}

// NewUserRepository creates a new instance of userRepositoryImpl.
func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		users:  make(map[int]models.User),
		lastID: 0,
	}
}

// CreateUser adds a new user to the in-memory database.
func (repo *userRepositoryImpl) CreateUser(user models.User) models.User {
	repo.lastID++
	user.ID = repo.lastID
	repo.users[user.ID] = user
	return user
}

// GetAllUsers retrieves all users from the in-memory database.
func (repo *userRepositoryImpl) GetAllUsers() []models.User {
	var result []models.User
	for _, user := range repo.users {
		result = append(result, user)
	}
	return result
}

// GetUserByID retrieves a user by their ID.
func (repo *userRepositoryImpl) GetUserByID(id int) (models.User, bool) {
	user, exists := repo.users[id]
	return user, exists
}

// UpdateUser updates an existing user in the in-memory database.
func (repo *userRepositoryImpl) UpdateUser(user models.User) (models.User, bool) {
	_, exists := repo.users[user.ID]
	if exists {
		repo.users[user.ID] = user
		return user, true
	}
	return user, false
}

// DeleteUser removes a user from the in-memory database.
func (repo *userRepositoryImpl) DeleteUser(id int) bool {
	_, exists := repo.users[id]
	if exists {
		delete(repo.users, id)
		return true
	}
	return false
}
