package database

import (
	"time"

	"github.com/data-market/internal/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var logger = logs.Logger("database")

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("market.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to open database %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("failed to connect database %s", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	err = sqlDB.Ping()
	if err != nil {
		logger.Errorf("failed to ping database %s", err)
	}

	DB = db
}
