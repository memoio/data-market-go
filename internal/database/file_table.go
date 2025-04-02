package database

import (
	"time"
)

type File struct {
	FileID         uint       `gorm:"primaryKey;autoIncrement;column:file_id" form:"file_id" json:"file_id"`
	Name           string     `gorm:"not null;column:name" form:"name" json:"name"`
	Description    string     `gorm:"column:description" form:"description" json:"description"`
	FileType       string     `gorm:"not null;column:file_type" form:"file_type" json:"file_type"`
	Category       string     `gorm:"not null;column:category" form:"category" json:"category"`
	OwnerAddress   string     `gorm:"not null;index:idx_owner;column:owner_address" form:"owner_address" json:"owner_address"`
	UploadTime     time.Time  `gorm:"default:CURRENT_TIMESTAMP;column:upload_time" form:"upload_time" json:"upload_time"`
	PublishState   int        `gorm:"default:0;index:idx_publish_state;column:publish_state" form:"publish_state" json:"publish_state"`
	PublishTime    *time.Time `gorm:"column:publish_time" form:"publish_time" json:"publish_time"`
	Price          string     `gorm:"default:'0';column:price" form:"price" json:"price"`
	FileSize       int64      `gorm:"not null;column:file_size" form:"file_size" json:"file_size"`
	PurchaseCount  int        `gorm:"default:0;column:purchase_count" form:"purchase_count" json:"purchase_count"`
	DownloadCount  int        `gorm:"default:0;column:download_count" form:"download_count" json:"download_count"`
	ViewCount      int        `gorm:"default:0;column:view_count" form:"view_count" json:"view_count"`
	ETag           string     `gorm:"column:e_tag" form:"e_tag" json:"e_tag"`
	FileDID        string     `gorm:"index:idx_file_did;column:file_did" form:"file_did" json:"file_did"`
	FileDIDTopic   string     `gorm:"index:idx_file_did_topic;column:file_did_topic" form:"file_did_topic" json:"file_did_topic"`
	ControllerDID  string     `gorm:"index:idx_controller_did;column:controller_did" form:"controller_did" json:"controller_did"`
	ControllerAddr string     `gorm:"index:idx_controller_addr;column:controller_addr" form:"controller_addr" json:"controller_addr"`

	// 组合索引
	IndexFileTypeCategory string `gorm:"index:idx_type_category,priority:1;column:index_file_type_category" form:"index_file_type_category" json:"index_file_type_category"`
	IndexFileType         string `gorm:"index:idx_type_category,priority:2;column:index_file_type" form:"index_file_type" json:"index_file_type"`
}

func (File) TableName() string {
	return "file_info"
}
