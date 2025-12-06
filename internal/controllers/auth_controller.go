package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"lottery/internal/entities"
	"lottery/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthController handles authentication-related HTTP requests
type AuthController struct{}

// NewAuthController creates a new instance of AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RefreshRequest represents the request body for token refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// LogoutRequest represents the request body for logout
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
	User         AuthUserResponse `json:"user"`
}

// AuthUserResponse represents user data in authentication responses
type AuthUserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// Register handles POST /api/auth/register
// Creates a new user account and returns access + refresh tokens
func (ac *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Validate email format
	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid email format",
		})
		return
	}

	// Validate password length
	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: "Password must be at least 6 characters long",
		})
		return
	}

	// Check if user already exists
	var existingUser entities.User
	result := database.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		// User exists
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   "user_exists",
			Message: "User with this email already exists",
		})
		return
	}
	// If error is not "record not found", it's a database error
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Database error: " + result.Error.Error(),
		})
		return
	}
	// User doesn't exist (ErrRecordNotFound) - this is what we want, continue

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to hash password",
		})
		return
	}

	// Create user
	user := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		IsActive: true,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to create user",
		})
		return
	}

	// Generate tokens
	accessToken, err := generateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate access token",
		})
		return
	}

	refreshToken, err := generateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate refresh token",
		})
		return
	}

	// Save refresh token to database
	refreshTokenEntity := entities.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().AddDate(0, 0, getRefreshTokenExpirationDays()),
		Revoked:   false,
	}

	if err := database.DB.Create(&refreshTokenEntity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to save refresh token",
		})
		return
	}

	c.JSON(http.StatusCreated, AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: AuthUserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			IsActive: user.IsActive,
		},
	})
}

// Login handles POST /api/auth/login
// Authenticates user and returns access (15m) + refresh (7d) tokens
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Find user by email
	var user entities.User
	result := database.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		// Check if it's a "record not found" error
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Error:   "invalid_credentials",
				Message: "Invalid email or password",
			})
			return
		}
		// Other database errors
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Database error",
		})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusForbidden, ErrorResponse{
			Error:   "account_disabled",
			Message: "Account is disabled",
		})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Invalid email or password",
		})
		return
	}

	// Generate tokens
	accessToken, err := generateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate access token",
		})
		return
	}

	refreshToken, err := generateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate refresh token",
		})
		return
	}

	// Save refresh token to database
	refreshTokenEntity := entities.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().AddDate(0, 0, getRefreshTokenExpirationDays()),
		Revoked:   false,
	}

	if err := database.DB.Create(&refreshTokenEntity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to save refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: AuthUserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			IsActive: user.IsActive,
		},
	})
}

// Refresh handles POST /api/auth/refresh
// Validates refresh token and issues new access token
func (ac *AuthController) Refresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Find refresh token in database
	var refreshToken entities.RefreshToken
	result := database.DB.Where("token = ?", req.RefreshToken).Preload("User").First(&refreshToken)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_token",
			Message: "Invalid refresh token",
		})
		return
	}

	// Check if token is valid (not revoked and not expired)
	if !refreshToken.IsValid() {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_token",
			Message: "Refresh token is expired or revoked",
		})
		return
	}

	// Generate new access token
	accessToken, err := generateAccessToken(refreshToken.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to generate access token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

// Logout handles POST /api/auth/logout
// Marks refresh token as revoked
func (ac *AuthController) Logout(c *gin.Context) {
	var req LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	// Find and revoke refresh token
	result := database.DB.Model(&entities.RefreshToken{}).
		Where("token = ?", req.RefreshToken).
		Update("revoked", true)

	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "token_not_found",
			Message: "Refresh token not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// Helper functions

// generateAccessToken generates a JWT access token (15 minutes expiration)
func generateAccessToken(userID uint) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set. Please set it in your .env file")
	}

	if len(secret) < 32 {
		return "", errors.New("JWT_SECRET must be at least 32 characters long for security")
	}

	expirationMinutes := getAccessTokenExpirationMinutes()
	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// generateRefreshToken generates a random refresh token string
func generateRefreshToken(userID uint) (string, error) {
	// Generate 32 random bytes
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Encode to base64 URL-safe string
	token := base64.URLEncoding.EncodeToString(bytes)
	return token, nil
}

// getAccessTokenExpirationMinutes returns access token expiration in minutes
func getAccessTokenExpirationMinutes() int {
	expStr := os.Getenv("JWT_ACCESS_EXPIRATION_MINUTES")
	if expStr == "" {
		return 15 // default 15 minutes
	}
	exp, err := strconv.Atoi(expStr)
	if err != nil {
		return 15
	}
	return exp
}

// getRefreshTokenExpirationDays returns refresh token expiration in days
func getRefreshTokenExpirationDays() int {
	expStr := os.Getenv("JWT_REFRESH_EXPIRATION_DAYS")
	if expStr == "" {
		return 7 // default 7 days
	}
	exp, err := strconv.Atoi(expStr)
	if err != nil {
		return 7
	}
	return exp
}

// isValidEmail validates email format using regex
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
