package models

type Restaurant struct {
	ID        uint   `gorm:"autoIncrement;primaryKey"`
	Name      string `gorm:"not null"`
	Address   string `gorm:"not null"`
	Phone     int    `gorm:"not null"`
	IsOnline  *bool  `gorm:"default:true"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	// Relations
	Tables []Table // OnetoMany
	Items  []Item  // OneToMany
}
