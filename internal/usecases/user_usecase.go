package usecases

import (
	// TODO: Add imports when implementing
	// "errors"
	// "lottery/internal/entities"
	// "lottery/internal/repositories"
)

// UserRepository defines the interface for user repository
// This should match repositories.UserRepository interface
type UserRepository interface {
	Create(user interface{}) error
	FindByID(id uint) (interface{}, error)
	FindByUsername(username string) (interface{}, error)
	FindByEmail(email string) (interface{}, error)
	Update(user interface{}) error
	Delete(id uint) error
}

// userUseCase implements UserUseCase interface
type userUseCase struct {
	// userRepo is the repository for user data access
	userRepo UserRepository
}

// NewUserUseCase creates a new instance of UserUseCase
// It should receive the user repository as a dependency
func NewUserUseCase(userRepo UserRepository) UserUseCase {
	// TODO: Return initialized userUseCase with the repository
	return nil
}

// RegisterUser implements UserUseCase.RegisterUser
func (uc *userUseCase) RegisterUser(username, email, password string) (*UserResponse, error) {
	// TODO: Implement user registration logic:
	// 1. Validate input (non-empty, valid email format, password strength)
	// 2. Check if username already exists (call userRepo.FindByUsername)
	// 3. Check if email already exists (call userRepo.FindByEmail)
	// 4. Hash password using bcrypt or similar
	// 5. Create user entity with hashed password
	// 6. Save user (call userRepo.Create)
	// 7. Convert entity to UserResponse (without password hash)
	// 8. Return UserResponse
	return nil, nil
}

// GetUserByID implements UserUseCase.GetUserByID
func (uc *userUseCase) GetUserByID(userID uint) (*UserResponse, error) {
	// TODO: Implement get user by ID:
	// 1. Call userRepo.FindByID(userID)
	// 2. If not found, return error
	// 3. Convert entity to UserResponse
	// 4. Return UserResponse
	return nil, nil
}

