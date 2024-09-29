package util

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator"
)

type User struct {
	Username string   `json:"username" validate:"required"`
	Password string   `json:"-"`
	Email    string   `json:"email" validate:"required,email"`
	Age      int      `json:"age" validate:"required,min=18,max=99"`
	Hobbies  []string `json:"hobbies,omitempty"`
}

func JSONTag() {
	user := User{
		Username: "admin",
		Password: "root",
		Email:    "admin@email.com",
		Age:      22,
	}

	b, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	//Validate JSON
	err = validator.New().Struct(user)
	if err != nil {
		log.Fatalf("Validation failed due to %v\n", err)
	}

	fmt.Println(string(b))
}
