package server

import "time"

// 统一响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TransactionRecord struct {
	FileName     string    `json:"fileName"`
	Description  string    `json:"description"`
	Price        string    `json:"price"`
	BuyTime      time.Time `json:"buyTime"`
	BuyerAddress string    `json:"buyerAddress"`
}

// File 对应数据库表结构（添加JSON标签）
type File struct {
	ID                    uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                  string     `gorm:"not null" json:"name"`
	Description           string     `json:"description"`
	FileType              string     `gorm:"not null" json:"fileType"`
	Category              string     `gorm:"not null" json:"category"`
	OwnerAddress          string     `gorm:"not null;index:idx_owner" json:"ownerAddress"`
	UploadTime            time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"uploadTime"`
	PublishState          int        `gorm:"not null;default:0;index:idx_publish_state" json:"publishState"`
	PublishTime           *time.Time `json:"publishTime,omitempty"`
	Price                 string     `gorm:"not null;default:'0'" json:"price"`
	FileSize              int64      `gorm:"not null" json:"fileSize"`
	PurchaseCount         int        `gorm:"default:0" json:"purchaseCount"`
	DownloadCount         int        `gorm:"default:0" json:"downloadCount"`
	ViewCount             int        `gorm:"default:0" json:"viewCount"`
	ETag                  string     `gorm:"unique" json:"eTag,omitempty"`
	FileDID               string     `gorm:"unique;not null;index:idx_file_did" json:"fileDID"`
	IndexFileTypeCategory string     `gorm:"index:idx_type_category,priority:1" json:"-"`
	IndexFileType         string     `gorm:"index:idx_type_category,priority:2" json:"-"`
}
