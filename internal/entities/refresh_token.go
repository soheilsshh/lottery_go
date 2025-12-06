package entities

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model

	UserID uint `json:"user_id" gorm:"not null;index"`
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Token string `json:"token" gorm:"uniqueIndex;not null;size:500"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null;index"`
	Revoked bool `json:"revoked" gorm:"default:false;index"`
}


func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

func (rt *RefreshToken) IsValid() bool {
	return !rt.Revoked && !rt.IsExpired()
}
