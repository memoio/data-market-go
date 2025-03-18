package database

import (
	"fmt"
	"log"
	"time"

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

	// get sql db from gorm db
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置超时时间
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	// ping db
	logger.Debug("ping db")
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	// 检查数据库连接示例
	logger.Debug("checking db connection")
	if err := db.Raw("SELECT 1").Error; err != nil {
		log.Fatal("数据库连接失败: ", err)
	}

	// 启用外键约束（SQLite 默认关闭）
	_ = db.Exec("PRAGMA foreign_keys = ON")

	logger.Debug("migrating tables..")

	// 自动迁移表结构
	err = db.AutoMigrate(
		&BlockNumber{},
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
