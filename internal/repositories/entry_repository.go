package repositories

import (
	// TODO: Add imports when implementing
	// "errors"
	// "gorm.io/gorm"
	// "lottery/internal/entities"
)

// entryRepository implements EntryRepository interface
type entryRepository struct {
	// db is the database connection (GORM or raw SQL)
	db interface{} // TODO: Replace with *gorm.DB when implementing
}

// NewEntryRepository creates a new instance of EntryRepository
// It should receive the database connection as a dependency
func NewEntryRepository(db interface{}) EntryRepository {
	// TODO: Replace db parameter with *gorm.DB
	// TODO: Return initialized entryRepository with db connection
	return nil
}

// Create implements EntryRepository.Create
func (r *entryRepository) Create(entry interface{}) error {
	// TODO: Parameter should be *entities.Entry
	// TODO: Implement create entry:
	// - Use GORM: db.Create(entry)
	// - Handle errors (duplicate entry, foreign key constraints)
	// - Return error if creation fails
	return nil
}

// FindByID implements EntryRepository.FindByID
func (r *entryRepository) FindByID(id uint) (interface{}, error) {
	// TODO: Return type should be (*entities.Entry, error)
	// TODO: Implement find by ID:
	// - Use GORM: db.Preload("User").Preload("Draw").First(&entry, id)
	// - Preload User and Draw relationships
	// - Return error if not found
	// - Return entry and nil error if found
	return nil, nil
}

// FindByDrawID implements EntryRepository.FindByDrawID
func (r *entryRepository) FindByDrawID(drawID uint) (interface{}, error) {
	// TODO: Return type should be ([]entities.Entry, error)
	// TODO: Implement find by draw ID:
	// - Use GORM: db.Where("draw_id = ?", drawID).Preload("User").Find(&entries)
	// - Preload User to include user information
	// - Return empty slice and nil error if no entries found
	// - Return entries slice and nil error
	return nil, nil
}

// FindByUserID implements EntryRepository.FindByUserID
func (r *entryRepository) FindByUserID(userID uint) (interface{}, error) {
	// TODO: Return type should be ([]entities.Entry, error)
	// TODO: Implement find by user ID:
	// - Use GORM: db.Where("user_id = ?", userID).Preload("Draw").Find(&entries)
	// - Preload Draw to include draw information
	// - Return empty slice and nil error if no entries found
	// - Return entries slice and nil error
	return nil, nil
}

// FindByUserAndDraw implements EntryRepository.FindByUserAndDraw
func (r *entryRepository) FindByUserAndDraw(userID, drawID uint) (interface{}, error) {
	// TODO: Return type should be (*entities.Entry, error)
	// TODO: Implement find by user and draw:
	// - Use GORM: db.Where("user_id = ? AND draw_id = ?", userID, drawID).First(&entry)
	// - Return error if not found (gorm.ErrRecordNotFound)
	// - Return entry and nil error if found
	// - This is used to check if user already entered a draw
	return nil, nil
}

// CountByDrawID implements EntryRepository.CountByDrawID
func (r *entryRepository) CountByDrawID(drawID uint) (int64, error) {
	// TODO: Implement count by draw ID:
	// - Use GORM: db.Model(&Entry{}).Where("draw_id = ?", drawID).Count(&count)
	// - Return count and nil error
	return 0, nil
}

