package models

import "time"

type Application struct {
	ID        uint      `gorm:"primaryKey"`
	JobID     uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relación con Users (N:N)
	Users []UserApplication `gorm:"foreignKey:ApplicationID"`
}