package models

import "time"

type Application struct {
	ID        uint      `gorm:"primaryKey"`
	JobID     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relación con User
	Users []User `gorm:"many2many:user_applications"`
}
