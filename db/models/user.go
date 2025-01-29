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
    FirstName string    `gorm:"size:50"`
    LastName  string    `gorm:"size:50"`
    Role      Role      `gorm:"size:20;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
    Profile   *UserProfile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
