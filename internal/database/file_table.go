package database

import (
	"time"
)

// 定义所有表模型
type File struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"not null"`
	Description   string
	FileType      string    `gorm:"not null"`
	Category      string    `gorm:"not null"`
	OwnerAddress  string    `gorm:"not null;index:idx_owner"`
	UploadTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	PublishState  int       `gorm:"not null;default:0;index:idx_publish_state"` // 发布状态：0-未上架，1-已上架，2-已下架
	PublishTime   *time.Time
	Price         string `gorm:"not null;default:'0'"`
	FileSize      int64  `gorm:"not null"`
	PurchaseCount int    `gorm:"default:0"`
	DownloadCount int    `gorm:"default:0"`
	ViewCount     int    `gorm:"default:0"`
	ETag          string `gorm:"unique"`
	FileDID       string `gorm:"unique;not null;index:idx_file_did"`

	// 组合索引
	IndexFileTypeCategory string `gorm:"index:idx_type_category,priority:1"`
	IndexFileType         string `gorm:"index:idx_type_category,priority:2"`
}

func (File) TableName() string {
	return "file_info"
}
