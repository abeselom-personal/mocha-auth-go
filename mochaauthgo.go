package mochaauthgo

import (
	"errors"
	"fmt"
)

// User represents a basic user structure
type User struct {
	Username string
	Password string
}

// In-memory user store (for testing, replace with DB)
var users = map[string]string{
	"admin": "password123",
	"user":  "pass1234",
}

// Authenticate checks if a username/password combination is correct
func Authenticate(username, password string) (bool, error) {
	if storedPassword, exists := users[username]; exists {
		if storedPassword == password {
			return true, nil
		}
		return false, errors.New("invalid password")
	}
	return false, errors.New("user not found")
}

// Example usage
func ExampleUsage() {
	success, err := Authenticate("admin", "password123")
	if err != nil {
		fmt.Println("Authentication failed:", err)
	} else {
		fmt.Println("Authentication success:", success)
	}
}
