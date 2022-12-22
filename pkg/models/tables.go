package models

type Table struct {
	ID        uint `gorm:"autoIncrement;primaryKey"`
	Number    uint `gorm:"not null"`
	Name      string
	Capacity  uint  `gorm:"not null"`
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
