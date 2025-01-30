package models

type UserApplication struct {
	UserID        uint `gorm:"primaryKey;not null"`
	ApplicationID uint `gorm:"primaryKey;not null"`

	// Relaciones
	User        User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Application Application `gorm:"foreignKey:ApplicationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
