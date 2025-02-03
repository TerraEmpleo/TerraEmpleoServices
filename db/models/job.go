package models

import "time"

type Job struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"size:200;not null"`
    Description string    `gorm:"type:text;not null"`
    Location    string    `gorm:"size:100"`
    Salary      float64   `gorm:"not null"`
    Requirements string   `gorm:"type:text"`
    EmployerID  uint      `gorm:"not null"`
    CategoryID  uint      `gorm:"not null"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`

    Employer User     `gorm:"foreignKey:EmployerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    Category Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
