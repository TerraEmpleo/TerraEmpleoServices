package models

import "time"

type UserProfile struct {
    ID         uint      `gorm:"primaryKey"`
    UserID     uint      `gorm:"not null;uniqueIndex"`
    Location   string    `gorm:"size:100"`
    Skills     string    `gorm:"type:text"`
    Experience int       `gorm:"not null"`
    ResumeURL  string    `gorm:"size:255"`
    Bio        string    `gorm:"type:text"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`

    // Relación con User
    User *User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}