package models



type UserApplication struct {
	UserID        uint   `json:"user_id" gorm:"primaryKey;not null"`
	ApplicationID uint   `json:"application_id" gorm:"primaryKey;not null"`
	Status        string `json:"status" gorm:"size:20;default:'pending'"` // "pending", "approved", "rejected"


	// Relaciones
	User        User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Application Application `gorm:"foreignKey:ApplicationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}