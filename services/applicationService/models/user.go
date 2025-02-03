package models

import "time"

type Role string

const (
    RoleAdmin   Role = "admin"
    RoleFarmer  Role = "farmer"
    RoleEmployer Role = "employer"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"size:100;uniqueIndex;not null"`
	Role      Role      `gorm:"size:20;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relación con Applications (N:N)
	Applications []Application `gorm:"many2many:user_applications"`
}