package database

import (
	"time"
)

type Collection struct {
	// 复合主键（用户地址 + 备忘录DID）
	UserAddress string    `gorm:"type:varchar(42);primaryKey"`              // 主键部分1
	MemoDID     string    `gorm:"type:varchar(255);primaryKey"`             // 主键部分2
	CollectTime time.Time `gorm:"not null;index;default:CURRENT_TIMESTAMP"` // 收藏时间
}

func (Collection) TableName() string {
	return "collections"
}
