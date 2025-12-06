package controllers

import (
	// TODO: Add imports when implementing
	// "net/http"
	// "strconv"
	// "github.com/gin-gonic/gin"
	// "lottery/internal/usecases"
)

// DrawUseCase defines the interface for draw use case
// This should match usecases.DrawUseCase interface
type DrawUseCase interface {
	CreateDraw(name string) (*DrawResponse, error)
	EnterDraw(userID, drawID uint) (*EntryResponse, error)
	PerformDraw(drawID uint) (*DrawResponse, error)
	GetPastWinners() ([]WinnerResponse, error)
	GetDrawHistory() ([]DrawResponse, error)
}

// DrawResponse represents draw data returned to clients
type DrawResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	WinnerID  *uint  `json:"winner_id,omitempty"`
	CreatedAt int64  `json:"created_at"`
	DrawnAt   *int64 `json:"drawn_at,omitempty"`
}

// EntryResponse represents entry data returned to clients
type EntryResponse struct {
	ID        uint  `json:"id"`
	UserID    uint  `json:"user_id"`
	DrawID    uint  `json:"draw_id"`
	CreatedAt int64 `json:"created_at"`
}

// WinnerResponse represents winner information for past winners list
type WinnerResponse struct {
	DrawID   uint   `json:"draw_id"`
	DrawName string `json:"draw_name"`
	WinnerID uint   `json:"winner_id"`
	DrawnAt  int64  `json:"drawn_at"`
}

// DrawController handles HTTP requests related to draws
type DrawController struct {
	// drawUseCase is the use case for draw business logic
	drawUseCase DrawUseCase
}

// NewDrawController creates a new instance of DrawController
// It should receive the draw use case as a dependency
func NewDrawController(drawUseCase DrawUseCase) *DrawController {
	// TODO: Return initialized DrawController with use case
	return nil
}

// CreateDraw handles POST /api/v1/draws
// Request body: { "name": "..." }
// Response: 201 Created with draw data
// TODO: May require admin authentication in the future
// Parameter: c *gin.Context
func (ctrl *DrawController) CreateDraw(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement create draw handler:
	// 1. Parse JSON request body (name)
	// 2. Validate input
	// 3. Call drawUseCase.CreateDraw
	// 4. Handle errors
	// 5. Return 201 Created with draw data (use HandleSuccess)
}

// EnterDraw handles POST /api/v1/draws/:id/enter
// Requires: JWT authentication (AuthMiddleware)
// Response: 201 Created with entry data
// Parameter: c *gin.Context
func (ctrl *DrawController) EnterDraw(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement enter draw handler:
	// 1. Get user ID from context (GetUserIDFromContext)
	// 2. Get draw ID from URL parameter (c.Param("id"))
	// 3. Convert draw ID to uint
	// 4. Call drawUseCase.EnterDraw(userID, drawID)
	// 5. Handle errors (draw not found, already entered, draw closed)
	// 6. Return 201 Created with entry data (use HandleSuccess)
}

// PerformDraw handles POST /api/v1/draws/:id/perform
// Response: 200 OK with draw data including winner
// TODO: May require admin authentication in the future
// Parameter: c *gin.Context
func (ctrl *DrawController) PerformDraw(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement perform draw handler:
	// 1. Get draw ID from URL parameter (c.Param("id"))
	// 2. Convert draw ID to uint
	// 3. Call drawUseCase.PerformDraw(drawID)
	// 4. Handle errors (draw not found, draw already closed, no entries)
	// 5. Return 200 OK with draw data including winner (use HandleSuccess)
}

// GetPastWinners handles GET /api/v1/draws/winners
// Response: 200 OK with list of past winners
// Parameter: c *gin.Context
func (ctrl *DrawController) GetPastWinners(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement get past winners handler:
	// 1. Call drawUseCase.GetPastWinners
	// 2. Handle errors
	// 3. Return 200 OK with winners list (use HandleSuccess)
}

// GetDrawHistory handles GET /api/v1/draws
// Response: 200 OK with list of all draws
// Parameter: c *gin.Context
func (ctrl *DrawController) GetDrawHistory(c interface{}) {
	// TODO: Parameter should be *gin.Context
	// TODO: Implement get draw history handler:
	// 1. Call drawUseCase.GetDrawHistory
	// 2. Handle errors
	// 3. Return 200 OK with draws list (use HandleSuccess)
}

