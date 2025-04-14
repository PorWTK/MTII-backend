package entities

import "time"

type Authorization struct {
	Token     string    `gorm:"type:varchar(255)" json:"token"`
	ExpiresAt time.Time `gorm:"type:date" json:"expires_at"`
}
