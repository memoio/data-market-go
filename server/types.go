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

// 更新响应结构体（包含所有关键字段）
type OwnerFileResponse struct {
	FileName      string     `json:"fileName"`
	FileDID       string     `json:"fileDID"` // 新增字段
	FileType      string     `json:"fileType"`
	Category      string     `json:"category"`
	Price         string     `json:"price"`
	FileSize      int64      `json:"fileSize"`
	UploadTime    time.Time  `json:"uploadTime"`
	PublishState  int        `json:"publishState"`
	PublishTime   *time.Time `json:"publishTime,omitempty"`
	PurchaseCount int        `json:"purchaseCount"`
	DownloadCount int        `json:"downloadCount"` // 新增字段
	ViewCount     int        `json:"viewCount"`     // 新增字段
	Description   string     `json:"description,omitempty"`
	ETag          string     `json:"eTag,omitempty"` // 新增字段
}

// 首先定义响应数据结构体
type FileDownloadResponse struct {
	FileName        string    `json:"fileName"`
	FileDescription string    `json:"fileDescription"`
	FilePrice       string    `json:"filePrice"`
	DownloadTime    time.Time `json:"downloadTime"`
}
