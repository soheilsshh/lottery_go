package controllers

import (
	// TODO: Add imports when implementing JWT
	// "net/http"
	// "github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a Gin middleware for JWT authentication
// It should:
// 1. Extract JWT token from Authorization header (Bearer token)
// 2. Validate the token
// 3. Extract user ID from token claims
// 4. Set user ID in context for use in handlers
// 5. Call next handler if valid, return 401 if invalid
func AuthMiddleware() interface{} {
	// TODO: Return type should be gin.HandlerFunc
	// TODO: Implement JWT authentication middleware
	// - Read Authorization header
	// - Parse Bearer token
	// - Validate JWT token
	// - Extract user ID from claims
	// - Set user ID in gin.Context
	// - Call c.Next() if valid
	// - Return 401 Unauthorized if invalid
	// Example:
	// return func(c *gin.Context) {
	//     // Implementation
	//     c.Next()
	// }
	return nil
}

// GetUserIDFromContext extracts the user ID from Gin context
// This should be called after AuthMiddleware has set the user ID
// Parameter: c *gin.Context
func GetUserIDFromContext(c interface{}) (uint, error) {
	// TODO: Parameter should be *gin.Context
	// TODO: Extract user ID from context
	// - Get user ID from c.Get("user_id") or similar
	// - Convert to uint
	// - Return error if not found
	return 0, nil
}

// ErrorResponse represents a standard error response format
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SuccessResponse represents a standard success response format
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// HandleError is a helper function to return error responses consistently
// Parameter: c *gin.Context
func HandleError(c interface{}, statusCode int, err error, message string) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement error response handler
	// - Create ErrorResponse struct
	// - Return JSON response with status code
}

// HandleSuccess is a helper function to return success responses consistently
// Parameter: c *gin.Context
func HandleSuccess(c interface{}, statusCode int, data interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement success response handler
	// - Create SuccessResponse struct
	// - Return JSON response with status code
}

