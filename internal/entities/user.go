package entities

import (
	"gorm.io/gorm"
)

// User نماینده کاربر در سیستم
// User represents a user entity in the system
type User struct {
	gorm.Model

	// Name نام کاربر
	// Name is the user's name
	Name string `json:"name" gorm:"not null"`

	// Email ایمیل کاربر (یکتا)
	// Email is the user's email address (unique)
	Email string `json:"email" gorm:"uniqueIndex;not null"`

	// Password رمز عبور هش شده
	// Password is the hashed password
	Password string `json:"-" gorm:"not null"`

	// IsActive وضعیت فعال بودن کاربر
	// IsActive indicates if the user is active
	IsActive bool `json:"is_active" gorm:"default:true"`
}

// TableName نام جدول در دیتابیس
// TableName returns the database table name for User
func (User) TableName() string {
	return "users"
}
