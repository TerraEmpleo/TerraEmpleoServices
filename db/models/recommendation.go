package models

import "time"

type Recommendation struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    JobID     uint      `gorm:"not null"`
    Score     float64   `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`

    User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    Job  Job  `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
