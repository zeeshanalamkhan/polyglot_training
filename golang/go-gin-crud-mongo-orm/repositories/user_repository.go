// Package repositories contains the repository layer for interacting with the in-memory database.
package repositories

import (
	"context"
	"log"

	"github.com/gitish/polyglot_training/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository defines the interface for user-related database operations.
type UserRepository interface {
	CreateUser(user models.User) models.User
	GetAllUsers() []models.User
	GetUserByID(id int) (models.User, bool)
	UpdateUser(user models.User) (models.User, bool)
	DeleteUser(id int) bool
}

// userRepositoryImpl implements the UserRepository interface using MongoDB.
type userRepositoryImpl struct {
	collection *mongo.Collection // MongoDB collection for users
}

// NewUserRepository creates a new instance of userRepositoryImpl.
func NewUserRepository(client *mongo.Client) UserRepository {

	collection := client.Database("users_db").Collection("users") // Specify your database name
	return &userRepositoryImpl{collection: collection}
}

// CreateUser adds a new user to the in-memory database.
func (repo *userRepositoryImpl) CreateUser(user models.User) models.User {
	result, err := repo.collection.InsertOne(context.Background(), user) // Insert user into MongoDB
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err) // Handle insertion error
	}
	user.ID = result.InsertedID.(int) // Set the ID for the created user
	return user
}

// GetAllUsers retrieves all users from the in-memory database.
func (repo *userRepositoryImpl) GetAllUsers() []models.User {
	cursor, err := repo.collection.Find(context.Background(), bson.D{}) // Find all users
	if err != nil {
		log.Fatalf("Failed to find users: %v", err) // Handle retrieval error
	}
	defer cursor.Close(context.Background()) // Ensure cursor is closed after use

	var users []models.User
	for cursor.Next(context.Background()) { // Iterate over the cursor
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatalf("Failed to decode user: %v", err) // Handle decoding error
		}
		users = append(users, user) // Append user to the result slice
	}
	return users // Return all users
}

// GetUserByID retrieves a user by their ID.
func (repo *userRepositoryImpl) GetUserByID(id int) (models.User, bool) {
	var user models.User
	err := repo.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user) // Find user by ID
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, false // Return false if no user is found
		}
		log.Fatalf("Failed to find user: %v", err) // Handle other errors
	}
	return user, true // Return found user
}

// UpdateUser updates an existing user in the in-memory database.
func (repo *userRepositoryImpl) UpdateUser(user models.User) (models.User, bool) {
	filter := bson.M{"id": user.ID}                                                // Filter for the user to update
	update := bson.M{"$set": user}                                                 // Update fields with the user data
	result, err := repo.collection.UpdateOne(context.Background(), filter, update) // Perform the update
	if err != nil || result.MatchedCount == 0 {
		return user, false // Return false if update fails or no documents match
	}
	return user, true // Return updated user
}

// DeleteUser removes a user from the in-memory database.
func (repo *userRepositoryImpl) DeleteUser(id int) bool {
	filter := bson.M{"id": id}                                             // Filter for the user to delete
	result, err := repo.collection.DeleteOne(context.Background(), filter) // Perform the deletion
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err) // Handle deletion error
	}
	return result.DeletedCount > 0 // Return true if a document was deleted

}
