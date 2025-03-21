package database

import (
	"time"
)

// 定义所有表模型
type File struct {
	FileID        uint       `gorm:"primaryKey;autoIncrement;column:file_id"`
	Name          string     `gorm:"not null;column:name"`
	Description   string     `gorm:"column:description"`
	FileType      string     `gorm:"not null;column:file_type"`
	Category      string     `gorm:"not null;column:category"`
	OwnerAddress  string     `gorm:"not null;index:idx_owner;column:owner_address"`
	UploadTime    time.Time  `gorm:"default:CURRENT_TIMESTAMP;column:upload_time"`
	PublishState  int        `gorm:"not null;default:0;index:idx_publish_state;column:publish_state"` // 发布状态：0-未上架，1-已上架，2-已下架
	PublishTime   *time.Time `gorm:"column:publish_time"`
	Price         string     `gorm:"not null;default:'0';column:price"`
	FileSize      int64      `gorm:"not null;column:file_size"`
	PurchaseCount int        `gorm:"default:0;column:purchase_count"`
	DownloadCount int        `gorm:"default:0;column:download_count"`
	ViewCount     int        `gorm:"default:0;column:view_count"`
	ETag          string     `gorm:"unique;column:e_tag"`
	FileDID       string     `gorm:"unique;not null;index:idx_file_did;column:file_did"`

	// 组合索引（注意：字段需要实际存储数据，否则索引无效）
	IndexFileTypeCategory string `gorm:"index:idx_type_category,priority:1;column:index_file_type_category"`
	IndexFileType         string `gorm:"index:idx_type_category,priority:2;column:index_file_type"`
}

func (File) TableName() string {
	return "file_info"
}
