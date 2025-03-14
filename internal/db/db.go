package db

import (
	"fmt"
	"time"

	"github.com/data-market/internal/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var logger = logs.Logger("db")

var G_DB *gorm.DB

// 定义所有表模型
type File struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"not null"`
	Description   string
	FileType      string    `gorm:"not null"`
	Category      string    `gorm:"not null"`
	OwnerAddress  string    `gorm:"not null;index:idx_owner"`
	UploadTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
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

type Download struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserAddress  string    `gorm:"not null"`
	MemoDID      string    `gorm:"not null;column:memo_did"`
	FileDID      string    `gorm:"not null;index"`
	DownloadDate time.Time `gorm:"default:CURRENT_TIMESTAMP;index:idx_download_date"`
}

func (Download) TableName() string {
	return "download_record"
}

type MemoDID struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	MemoDID     string `gorm:"column:memo_did;type:TEXT;not null"`
	UserAddress string `gorm:"column:user_address;type:TEXT;not null"`
}

func (MemoDID) TableName() string {
	return "memodid"
}

type FileMemo struct {
	FileDID string `gorm:"column:file_did;type:TEXT;not null;primaryKey"`
	MemoDID string `gorm:"column:memo_did;type:TEXT;not null;primaryKey"`
	AddType int    `gorm:"column:add_type;type:INTEGER;not null;check:add_type IN (1,2)"`
}

func (FileMemo) TableName() string {
	return "file_memo"
}

func init() {
	logger.Debug("init db..")

	// 连接 SQLite 数据库（文件名为 market.db）
	db, err := gorm.Open(sqlite.Open("market.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败: %v", err))
	}

	// 启用外键约束（SQLite 默认关闭）
	_ = db.Exec("PRAGMA foreign_keys = ON")

	logger.Debug("migrating tables..")

	// 自动迁移表结构
	err = db.AutoMigrate(
		&File{},
		&Purchase{},
		&Download{},
		&MemoDID{},
		&FileMemo{},
	)
	if err != nil {
		panic(fmt.Errorf("表创建失败: %v", err))
	}

	// 手动创建组合索引（GORM 自动迁移可能不会处理）
	db.Exec("CREATE INDEX IF NOT EXISTS idx_type_category ON file_info(file_type, category)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_price ON file_info(price)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_upload_time ON file_info(upload_time)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_memo_did ON memodid (memo_did)")

	// save to global db
	G_DB = db
}
