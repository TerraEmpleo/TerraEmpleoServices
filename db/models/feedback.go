package models

import "time"

type Feedback struct {
    ID             uint      `gorm:"primaryKey"`
    ReviewerID     uint      `gorm:"not null"`
    ReviewedUserID uint      `gorm:"not null"`
    Rating         int       `gorm:"not null"`
    Comments       string    `gorm:"type:text"`
    CreatedAt      time.Time `gorm:"autoCreateTime"`

    Reviewer     User `gorm:"foreignKey:ReviewerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    ReviewedUser User `gorm:"foreignKey:ReviewedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
