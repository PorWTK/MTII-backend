package entities

import (
	"time"
)

type User struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
}
