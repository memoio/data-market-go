package database

import (
	"fmt"

	"github.com/data-market/internal/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var logger = logs.Logger("db")

var G_DB *gorm.DB

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
