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
	// db, err := gorm.Open(sqlite.Open("market.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(fmt.Errorf("数据库连接失败: %v", err))
	// }

	// 关键参数配置
	dsn := "file:market.db?cache=shared&mode=rwc&_journal_mode=WAL&_busy_timeout=5000"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		// 推荐配置参数
		DisableForeignKeyConstraintWhenMigrating: true, // 避免外键约束问题
	})
	if err != nil {
		panic(err)
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
		&Download{},
		&MemoDID{},
		&Access{},
		&Collection{},
	)
	if err != nil {
		panic(fmt.Errorf("表创建失败: %v", err))
	}

	// 使用原生SQL执行语句插入一条初始记录到file表
	result := db.Exec(`
    INSERT INTO file_info (
        name, description, file_type, category, owner_address, 
        upload_time, price, file_size, e_tag, file_did, 
        file_did_topic, controller_did, publish_state, publish_time, 
        index_file_type_category, index_file_type
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		"sample_document.pdf",
		"This is a sample document for demonstration",
		"document",
		"education",
		"0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
		time.Now(),
		nil, // Price设为NULL
		1048576,
		"bafkreih6n5g5w4y6u7uvc4mh7jhjm7gidmkrbbpi7phyiyg54gplvngcpm",
		"did:mfile:mid:bafkreih6n5g5w4y6u7uvc4mh7jhjm7gidmkrbbpi7phyiyg54gplvngcpm",
		"0x53268dc74973e787e3926c6c551ffd6621e9f0eaca199db8c75afdc3cb86837c",
		nil, // ControllerDID设为NULL
		0,
		nil,
		"document_education",
		"document",
	)

	if result.Error != nil {
		panic(result.Error)
	}

	// 手动创建组合索引（GORM 自动迁移可能不会处理）
	db.Exec("CREATE INDEX IF NOT EXISTS idx_type_category ON file_info(file_type, category)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_price ON file_info(price)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_upload_time ON file_info(upload_time)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_memo_did ON memodid (memo_did)")

	// save to global db
	G_DB = db
}
