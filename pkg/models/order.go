package models

import (
	"errors"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type PayementMode struct {
	Cash       string `json:"cash"`
	DebitCard  string `json:"debit_card"`
	CreditCard string `json:"credit_card"`
}

type Order struct {
	ID            uuid.UUID    `gorm:"type:uuid;primaryKey"`
	IsPaymentDone bool         `gorm:"default:true"`
	BillAmount    int          `gorm:"not null"`
	GST           int          `gorm:"default:18"`
	PaidVia       PayementMode `gorm:"not null"`
	CreatedAt     int64        `gorm:"autoCreateTime"`
	// Relations
	Table Table `gorm:"foreignKey:TableRefer"` // BelongsTo
	Item  Item  `gorm:"foreignKey:ItemRefer"`  // BelongsTo
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID, err = uuid.NewUUID()
	if err != nil {
		err = errors.New("can't save order data")
	}
	return
}
