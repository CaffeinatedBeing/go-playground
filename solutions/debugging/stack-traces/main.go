package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
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

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email format")
	}
	user := &User{
		ID:        int(time.Now().UnixNano()),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		Data:      make([]byte, 1000),
	}
	s.users[user.ID] = user
	return user, nil
}

func (s *UserService) GetUser(id int) (*User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, name, email string) error {
	user, ok := s.users[id]
	if !ok {
		return fmt.Errorf("user with id %d not found", id)
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("invalid email format")
	}
	user.Name = name
	user.Email = email
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user with id %d not found", id)
	}
	delete(s.users, id)
	return nil
}

func (s *UserService) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	data, err := json.Marshal(s.users)
	if err != nil {
		return fmt.Errorf("failed to marshal users: %w", err)
	}
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func (s *UserService) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	var users map[int]*User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return fmt.Errorf("failed to decode users: %w", err)
	}
	s.users = users
	return nil
}

func (s *UserService) ProcessUserData(id int) error {
	user, ok := s.users[id]
	if !ok {
		return fmt.Errorf("user with id %d not found", id)
	}
	if user.Data == nil {
		return errors.New("user data is nil")
	}
	for i := 0; i < len(user.Data); i++ {
		user.Data[i] = byte(i % 256)
	}
	return nil
}

// recoverWithStack recovers from panics and prints a stack trace
func recoverWithStack() {
	if r := recover(); r != nil {
		fmt.Fprintf(os.Stderr, "Recovered from panic: %v\n", r)
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		fmt.Fprintf(os.Stderr, "Stack trace:\n%s\n", buf[:stackSize])
	}
}

func main() {
	defer recoverWithStack()
	service := NewUserService()

	// Create some users
	user1, err := service.CreateUser("John Doe", "john@example.com")
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	_, err = service.CreateUser("Jane Smith", "jane@example.com")
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// Update a user
	err = service.UpdateUser(user1.ID, "John Updated", "john.updated@example.com")
	if err != nil {
		fmt.Println("Error updating user:", err)
	}

	// Process user data
	err = service.ProcessUserData(user1.ID)
	if err != nil {
		fmt.Println("Error processing user data:", err)
	}

	// Save to file
	err = service.SaveToFile("users.json")
	if err != nil {
		fmt.Println("Error saving to file:", err)
	}

	// Try to get a non-existent user
	_, err = service.GetUser(999)
	if err != nil {
		fmt.Println("Error getting non-existent user:", err)
	}

	// Try to update a non-existent user
	err = service.UpdateUser(999, "Non Existent", "nonexistent@example.com")
	if err != nil {
		fmt.Println("Error updating non-existent user:", err)
	}

	// Try to process data for a non-existent user
	err = service.ProcessUserData(999)
	if err != nil {
		fmt.Println("Error processing data for non-existent user:", err)
	}

	// Try to delete a non-existent user
	err = service.DeleteUser(999)
	if err != nil {
		fmt.Println("Error deleting non-existent user:", err)
	}

	// Try to load from a non-existent file
	err = service.LoadFromFile("nonexistent.json")
	if err != nil {
		fmt.Println("Error loading from non-existent file:", err)
	}
}
