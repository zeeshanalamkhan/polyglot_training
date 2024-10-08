// Package models contains data models for the application.
package models

// User represents a user in the system.
type User struct {
	ID    int    `bson:"id" json:"id"`       // Unique identifier for the user
	Name  string `bson:"name" json:"name"`   // User's name
	Email string `bson:"email" json:"email"` // User's email
}
