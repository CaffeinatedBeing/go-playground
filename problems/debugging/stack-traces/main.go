package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Data      []byte    `json:"data"`
}

// UserService handles user-related operations
type UserService struct {
	users map[int]*User
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(name, email string) *User {
	// BUG: No validation of input parameters
	user := &User{
		ID:        rand.Intn(1000),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		Data:      make([]byte, 1000),
	}
	s.users[user.ID] = user
	return user
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) *User {
	// BUG: No error handling for non-existent user
	return s.users[id]
}

// UpdateUser updates a user's information
func (s *UserService) UpdateUser(id int, name, email string) {
	// BUG: No error handling for non-existent user
	user := s.users[id]
	user.Name = name
	user.Email = email
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) {
	// BUG: No error handling for non-existent user
	delete(s.users, id)
}

// SaveToFile saves all users to a file
func (s *UserService) SaveToFile(filename string) error {
	// BUG: No error handling for file operations
	file, _ := os.Create(filename)
	defer file.Close()

	// BUG: No error handling for JSON marshaling
	data, _ := json.Marshal(s.users)
	file.Write(data)
	return nil
}

// LoadFromFile loads users from a file
func (s *UserService) LoadFromFile(filename string) error {
	// BUG: No error handling for file operations
	file, _ := os.Open(filename)
	defer file.Close()

	// BUG: No error handling for JSON unmarshaling
	var users map[int]*User
	json.NewDecoder(file).Decode(&users)
	s.users = users
	return nil
}

// ProcessUserData processes user data
func (s *UserService) ProcessUserData(id int) {
	// BUG: No error handling for non-existent user
	user := s.users[id]

	// BUG: Potential panic if user.Data is nil
	for i := 0; i < len(user.Data); i++ {
		user.Data[i] = byte(i % 256)
	}
}

func main() {
	service := NewUserService()

	// Create some users
	user1 := service.CreateUser("John Doe", "john@example.com")
	_ = service.CreateUser("Jane Smith", "jane@example.com")

	// Update a user
	service.UpdateUser(user1.ID, "John Updated", "john.updated@example.com")

	// Process user data
	service.ProcessUserData(user1.ID)

	// Save to file
	service.SaveToFile("users.json")

	// Try to get a non-existent user
	nonExistentUser := service.GetUser(999)
	fmt.Printf("Non-existent user: %+v\n", nonExistentUser)

	// Try to update a non-existent user
	service.UpdateUser(999, "Non Existent", "nonexistent@example.com")

	// Try to process data for a non-existent user
	service.ProcessUserData(999)

	// Try to delete a non-existent user
	service.DeleteUser(999)

	// Try to load from a non-existent file
	service.LoadFromFile("nonexistent.json")
}
