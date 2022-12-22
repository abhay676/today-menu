package models

type Item struct {
	ID          uint   `gorm:"autoIncrement;primaryKey"`
	Name        string `gorm:"not null"`
	Price       int16  `gorm:"not null"`
	Image       string
	Description string
	IsAvailable *bool `gorm:"default:true"`
	CreatedAt   int64 `gorm:"autoCreateTime"`
	UpdatedAt   int64 `gorm:"autoUpdateTime"`
}
