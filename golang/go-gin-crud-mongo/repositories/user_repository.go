package repositories

import (
	"context"
	"log"

	"github.com/gitish/polyglot_training/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() []models.User
	GetUserByID(id primitive.ObjectID) (models.User, bool)
	UpdateUser(user models.User) (models.User, bool)
	DeleteUser(id primitive.ObjectID) bool
}

// userRepositoryImpl implements the UserRepository interface using MongoDB.
type userRepositoryImpl struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new instance of userRepositoryImpl.
func NewUserRepository(client *mongo.Client) UserRepository {
	collection := client.Database("users_db").Collection("users")
	return &userRepositoryImpl{collection: collection}
}

// CreateUser adds a new user to the database.
func (repo *userRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	result, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
		return models.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

// GetAllUsers retrieves all users from the database.
func (repo *userRepositoryImpl) GetAllUsers() []models.User {
	cursor, err := repo.collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to find users: %v", err)
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Fatalf("Failed to decode user: %v", err)
		}
		users = append(users, user)
	}
	return users
}

// GetUserByID retrieves a user by their ID.
func (repo *userRepositoryImpl) GetUserByID(id primitive.ObjectID) (models.User, bool) {
	var user models.User
	err := repo.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, false
		}
		log.Fatalf("Failed to find user: %v", err)
	}
	return user, true
}

// UpdateUser updates an existing user in the database.
func (repo *userRepositoryImpl) UpdateUser(user models.User) (models.User, bool) {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	result, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
		return models.User{}, false
	}
	if result.MatchedCount == 0 {
		return models.User{}, false
	}
	return user, true
}

// DeleteUser deletes a user from the database by their ID.
func (repo *userRepositoryImpl) DeleteUser(id primitive.ObjectID) bool {
	filter := bson.M{"_id": id}
	result, err := repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
		return false
	}
	return result.DeletedCount > 0
}
