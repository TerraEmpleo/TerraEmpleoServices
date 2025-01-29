package models

import "time"

type Category struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"size:100;not null"`
    Description string    `gorm:"type:text"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
}
