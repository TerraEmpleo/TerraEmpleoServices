package models

import "time"

type Application struct {
    ID             uint      `gorm:"primaryKey"`
    JobID          uint      `gorm:"not null"`
    UserID         uint      `gorm:"not null"`
    Status         string    `gorm:"size:20"`
    ApplicationDate time.Time `gorm:"not null"`
    UpdatedAt      time.Time `gorm:"autoUpdateTime"`

    Job  Job  `gorm:"foreignKey:JobID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
