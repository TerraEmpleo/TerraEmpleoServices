package models


type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:50;uniqueIndex;not null"`
	Email    string `gorm:"size:100;uniqueIndex;not null"`
	Role     string `gorm:"size:20;not null"` // admin, farmer, employer
}