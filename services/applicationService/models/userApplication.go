package models

type UserApplication struct {
	UserID        uint `gorm:"primaryKey;not null"`
	ApplicationID uint `gorm:"primaryKey;not null"`
	Status        string `gorm:"size:20;default:'pending'"` // "pending", "approved", "rejected"

	// Relaciones
	User        *User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Application *Application `gorm:"foreignKey:ApplicationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
