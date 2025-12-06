package controllers

import (
	// TODO: Add imports when implementing
	// "net/http"
	// "github.com/gin-gonic/gin"
	// "lottery/internal/usecases"
)

// UserUseCase defines the interface for user use case
// This should match usecases.UserUseCase interface
type UserUseCase interface {
	RegisterUser(username, email, password string) (*UserResponse, error)
	GetUserByID(userID uint) (*UserResponse, error)
}

// UserResponse represents user data returned to clients
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

// UserController handles HTTP requests related to users
type UserController struct {
	// userUseCase is the use case for user business logic
	userUseCase UserUseCase
}

// NewUserController creates a new instance of UserController
// It should receive the user use case as a dependency
func NewUserController(userUseCase UserUseCase) *UserController {
	// TODO: Return initialized UserController with use case
	return nil
}

// Register handles POST /api/v1/users/register
// Request body: { "username": "...", "email": "...", "password": "..." }
// Response: 201 Created with user data (without password)
// Parameter: c *gin.Context
func (ctrl *UserController) Register(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement register handler:
	// 1. Parse JSON request body (username, email, password)
	// 2. Validate input
	// 3. Call userUseCase.RegisterUser
	// 4. Handle errors (duplicate username/email, validation errors)
	// 5. Return 201 Created with user data (use HandleSuccess)
}

// Login handles POST /api/v1/users/login
// Request body: { "username": "...", "password": "..." }
// Response: 200 OK with JWT token and user data
// TODO: Implement when JWT is ready
// Parameter: c *gin.Context
func (ctrl *UserController) Login(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement login handler:
	// 1. Parse JSON request body (username, password)
	// 2. Call userUseCase.LoginUser
	// 3. Handle errors (invalid credentials)
	// 4. Return 200 OK with JWT token and user data
}

// GetMe handles GET /api/v1/users/me
// Requires: JWT authentication (AuthMiddleware)
// Response: 200 OK with current user data
// Parameter: c *gin.Context
func (ctrl *UserController) GetMe(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement get me handler:
	// 1. Get user ID from context (GetUserIDFromContext)
	// 2. Call userUseCase.GetUserByID
	// 3. Handle errors (user not found)
	// 4. Return 200 OK with user data (use HandleSuccess)
}

