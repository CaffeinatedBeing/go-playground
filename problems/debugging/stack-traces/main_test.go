package main

import (
	"os"
	"testing"
)

func TestUserService(t *testing.T) {
	service := NewUserService()

	// Test creating a user
	user := service.CreateUser("Test User", "test@example.com")
	if user == nil {
		t.Fatal("CreateUser returned nil")
	}
	if user.Name != "Test User" {
		t.Errorf("Expected name 'Test User', got '%s'", user.Name)
	}

	// Test getting a user
	retrievedUser := service.GetUser(user.ID)
	if retrievedUser == nil {
		t.Fatal("GetUser returned nil for existing user")
	}
	if retrievedUser.ID != user.ID {
		t.Errorf("Expected ID %d, got %d", user.ID, retrievedUser.ID)
	}

	// Test updating a user
	newName := "Updated User"
	newEmail := "updated@example.com"
	service.UpdateUser(user.ID, newName, newEmail)
	updatedUser := service.GetUser(user.ID)
	if updatedUser.Name != newName {
		t.Errorf("Expected name '%s', got '%s'", newName, updatedUser.Name)
	}
	if updatedUser.Email != newEmail {
		t.Errorf("Expected email '%s', got '%s'", newEmail, updatedUser.Email)
	}

	// Test processing user data
	service.ProcessUserData(user.ID)
	processedUser := service.GetUser(user.ID)
	if len(processedUser.Data) != 1000 {
		t.Errorf("Expected data length 1000, got %d", len(processedUser.Data))
	}

	// Test file operations
	filename := "test_users.json"
	err := service.SaveToFile(filename)
	if err != nil {
		t.Errorf("SaveToFile failed: %v", err)
	}

	// Clean up the test file
	defer os.Remove(filename)

	// Test loading from file
	newService := NewUserService()
	err = newService.LoadFromFile(filename)
	if err != nil {
		t.Errorf("LoadFromFile failed: %v", err)
	}

	// Verify loaded data
	loadedUser := newService.GetUser(user.ID)
	if loadedUser == nil {
		t.Fatal("Failed to load user from file")
	}
	if loadedUser.Name != newName {
		t.Errorf("Expected name '%s', got '%s'", newName, loadedUser.Name)
	}
}

func TestErrorCases(t *testing.T) {
	service := NewUserService()

	// Test getting non-existent user
	nonExistentUser := service.GetUser(999)
	if nonExistentUser != nil {
		t.Error("GetUser should return nil for non-existent user")
	}

	// Test updating non-existent user
	// This should not panic
	service.UpdateUser(999, "Non Existent", "nonexistent@example.com")

	// Test processing data for non-existent user
	// This should not panic
	service.ProcessUserData(999)

	// Test deleting non-existent user
	// This should not panic
	service.DeleteUser(999)

	// Test loading from non-existent file
	err := service.LoadFromFile("nonexistent.json")
	if err == nil {
		t.Error("LoadFromFile should return error for non-existent file")
	}
}

func TestInputValidation(t *testing.T) {
	service := NewUserService()

	// Test creating user with empty name
	user := service.CreateUser("", "test@example.com")
	if user != nil {
		t.Error("CreateUser should not accept empty name")
	}

	// Test creating user with invalid email
	user = service.CreateUser("Test User", "invalid-email")
	if user != nil {
		t.Error("CreateUser should not accept invalid email")
	}
}
