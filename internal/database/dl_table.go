package database

import (
	"time"
)

type Download struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserAddress  string    `gorm:"not null;index:idx_user_address"` // 新增索引
	FileID       uint      `gorm:"not null;index:idx_file_id"`
	MemoDID      string    `gorm:"not null;column:memo_did"`
	FileDID      string    `gorm:"not null;index"`
	DownloadDate time.Time `gorm:"default:CURRENT_TIMESTAMP;index:idx_download_date"`
}

func (Download) TableName() string {
	return "downloads"
}
