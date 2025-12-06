package repositories

import (
	// TODO: Import entities when implementing
	// "lottery/internal/entities"
)

// UserRepository defines the interface for user data access
// Note: User type should be imported from entities package
type UserRepository interface {
	// Create saves a new user to the database
	// Parameter: user *entities.User
	Create(user interface{}) error

	// FindByID retrieves a user by their ID
	// Returns: *entities.User
	FindByID(id uint) (interface{}, error)

	// FindByUsername retrieves a user by their username
	// Returns: *entities.User
	FindByUsername(username string) (interface{}, error)

	// FindByEmail retrieves a user by their email
	// Returns: *entities.User
	FindByEmail(email string) (interface{}, error)

	// Update updates an existing user in the database
	// Parameter: user *entities.User
	Update(user interface{}) error

	// Delete removes a user from the database
	Delete(id uint) error
}

// DrawRepository defines the interface for draw data access
// Note: Draw type should be imported from entities package
type DrawRepository interface {
	// Create saves a new draw to the database
	// Parameter: draw *entities.Draw
	Create(draw interface{}) error

	// FindByID retrieves a draw by its ID
	// Returns: *entities.Draw
	FindByID(id uint) (interface{}, error)

	// FindAll retrieves all draws
	// Returns: []entities.Draw
	FindAll() (interface{}, error)

	// FindOpen retrieves all open draws
	// Returns: []entities.Draw
	FindOpen() (interface{}, error)

	// FindClosedWithWinners retrieves all closed draws that have winners
	// Returns: []entities.Draw
	FindClosedWithWinners() (interface{}, error)

	// Update updates an existing draw in the database
	// Parameter: draw *entities.Draw
	Update(draw interface{}) error

	// Delete removes a draw from the database
	Delete(id uint) error
}

// EntryRepository defines the interface for entry data access
// Note: Entry type should be imported from entities package
type EntryRepository interface {
	// Create saves a new entry to the database
	// Parameter: entry *entities.Entry
	Create(entry interface{}) error

	// FindByID retrieves an entry by its ID
	// Returns: *entities.Entry
	FindByID(id uint) (interface{}, error)

	// FindByDrawID retrieves all entries for a specific draw
	// Returns: []entities.Entry
	FindByDrawID(drawID uint) (interface{}, error)

	// FindByUserID retrieves all entries for a specific user
	// Returns: []entities.Entry
	FindByUserID(userID uint) (interface{}, error)

	// FindByUserAndDraw checks if a user has already entered a specific draw
	// Returns: *entities.Entry
	FindByUserAndDraw(userID, drawID uint) (interface{}, error)

	// CountByDrawID counts the number of entries for a specific draw
	CountByDrawID(drawID uint) (int64, error)
}

