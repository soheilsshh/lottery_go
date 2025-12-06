package usecases

import (
	// TODO: Import entities when implementing
	// "lottery/internal/entities"
)

// UserUseCase defines the interface for user-related business logic
type UserUseCase interface {
	// RegisterUser creates a new user account
	// It should:
	// - Validate input (username, email, password)
	// - Check if username/email already exists
	// - Hash the password
	// - Save user to database
	// - Return the created user (without password hash)
	RegisterUser(username, email, password string) (*UserResponse, error)

	// LoginUser authenticates a user and returns a JWT token
	// It should:
	// - Validate credentials
	// - Check password hash
	// - Generate JWT token
	// - Return token and user info
	// TODO: Implement when JWT is ready
	// LoginUser(username, password string) (*LoginResponse, error)

	// GetUserByID retrieves a user by their ID
	GetUserByID(userID uint) (*UserResponse, error)
}

// DrawUseCase defines the interface for draw-related business logic
type DrawUseCase interface {
	// CreateDraw creates a new lottery draw
	// It should:
	// - Validate input
	// - Set status to "open"
	// - Save to database
	CreateDraw(name string) (*DrawResponse, error)

	// EnterDraw allows a user to enter a draw
	// It should:
	// - Check if draw exists and is open
	// - Check if user already entered this draw
	// - Create entry record
	// - Return entry info
	EnterDraw(userID, drawID uint) (*EntryResponse, error)

	// PerformDraw selects a random winner from all entries
	// It should:
	// - Check if draw exists and is open
	// - Get all entries for this draw
	// - Randomly select one entry
	// - Set winner and close the draw
	// - Return winner info
	PerformDraw(drawID uint) (*DrawResponse, error)

	// GetPastWinners returns a list of all past winners
	// It should:
	// - Query all closed draws with winners
	// - Include winner user information
	// - Return sorted list (newest first)
	GetPastWinners() ([]WinnerResponse, error)

	// GetDrawHistory returns all draws with their status
	GetDrawHistory() ([]DrawResponse, error)
}

// UserResponse represents user data returned to clients (without sensitive info)
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

// DrawResponse represents draw data returned to clients
type DrawResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Status    string  `json:"status"`
	WinnerID  *uint   `json:"winner_id,omitempty"`
	Winner    *UserResponse `json:"winner,omitempty"`
	CreatedAt int64   `json:"created_at"`
	DrawnAt   *int64  `json:"drawn_at,omitempty"`
}

// EntryResponse represents entry data returned to clients
type EntryResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	DrawID    uint   `json:"draw_id"`
	CreatedAt int64  `json:"created_at"`
}

// WinnerResponse represents winner information for past winners list
type WinnerResponse struct {
	DrawID    uint         `json:"draw_id"`
	DrawName  string       `json:"draw_name"`
	Winner    UserResponse `json:"winner"`
	DrawnAt   int64        `json:"drawn_at"`
}

