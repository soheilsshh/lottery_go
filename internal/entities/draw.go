package entities

import (
	"time"

	"gorm.io/gorm"
)

// Draw نماینده قرعه‌کشی در سیستم
// Draw represents a lottery draw entity
type Draw struct {
	gorm.Model

	// Title عنوان قرعه‌کشی
	// Title is the title of the draw
	Title string `json:"title" gorm:"not null"`

	// Description توضیحات قرعه‌کشی
	// Description is the description of the draw
	Description string `json:"description" gorm:"type:text"`

	// StartDate تاریخ شروع قرعه‌کشی
	// StartDate is the start date of the draw
	StartDate time.Time `json:"start_date" gorm:"not null"`

	// EndDate تاریخ پایان قرعه‌کشی
	// EndDate is the end date of the draw
	EndDate time.Time `json:"end_date" gorm:"not null"`

	// IsActive وضعیت فعال بودن قرعه‌کشی
	// IsActive indicates if the draw is active
	IsActive bool `json:"is_active" gorm:"default:true"`

	// DrawnAt زمان انجام قرعه‌کشی (null اگر هنوز انجام نشده)
	// DrawnAt is the time when the draw was performed (null if not drawn yet)
	DrawnAt *time.Time `json:"drawn_at,omitempty"`

	// WinnerID شناسه کاربر برنده (null اگر هنوز برنده انتخاب نشده)
	// WinnerID is the ID of the winning user (null if winner not selected yet)
	WinnerID *uint `json:"winner_id,omitempty" gorm:"index"`

	// Winner کاربر برنده (رابطه)
	// Winner is the user who won this draw (relationship)
	Winner *User `json:"winner,omitempty" gorm:"foreignKey:WinnerID"`
}

// TableName نام جدول در دیتابیس
// TableName returns the database table name for Draw
func (Draw) TableName() string {
	return "draws"
}

// DrawEntry نماینده ورود کاربر به یک قرعه‌کشی
// DrawEntry represents a user's entry into a draw
type DrawEntry struct {
	gorm.Model

	// UserID شناسه کاربر
	// UserID is the ID of the user who entered
	UserID uint `json:"user_id" gorm:"not null;index"`

	// User کاربری که وارد شده (رابطه - پیش‌بارگذاری)
	// User is the user who made this entry (relationship - preload)
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`

	// DrawID شناسه قرعه‌کشی
	// DrawID is the ID of the draw this entry belongs to
	DrawID uint `json:"draw_id" gorm:"not null;index"`

	// Draw قرعه‌کشی مربوطه (رابطه - پیش‌بارگذاری)
	// Draw is the draw this entry belongs to (relationship - preload)
	Draw Draw `json:"draw,omitempty" gorm:"foreignKey:DrawID"`
}

// TableName نام جدول در دیتابیس
// TableName returns the database table name for DrawEntry
func (DrawEntry) TableName() string {
	return "draw_entries"
}
