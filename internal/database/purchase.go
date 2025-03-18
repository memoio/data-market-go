package database

import (
	"time"
)

type Purchase struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserAddress  string    `gorm:"not null"`
	MemoDID      string    `gorm:"not null;column:memo_did"`
	FileDID      string    `gorm:"not null;index"`
	PurchaseDate time.Time `gorm:"default:CURRENT_TIMESTAMP;index:idx_purchase_date"`
	Price        string    `gorm:"not null"`
}

func (Purchase) TableName() string {
	return "purchase_record"
}
