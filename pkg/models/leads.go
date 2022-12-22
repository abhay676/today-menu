package models

type Leads struct {
	ID        string `gorm:"autoIncrement"`
	Number    int    `gorm:"primaryKey"`
	Email     string
	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
