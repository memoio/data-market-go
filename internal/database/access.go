package database

import "time"

// the access record of mfiles
type Access struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	FileDID      string    `gorm:"column:file_did;type:TEXT;not null"`
	FileID       uint      `gorm:"column:file_id;index;not null"` // 对应File表的ID
	MemoDID      string    `gorm:"column:memo_did;type:TEXT;not null"`
	UserAddress  string    `gorm:"column:user_address;type:TEXT;not null"`
	OwnerAddress string    `gorm:"column:owner_address;type:TEXT;not null"`
	AddTime      time.Time `gorm:"column:add_time;type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	AddType      int       `gorm:"column:add_type;type:INTEGER;not null;check:add_type IN (1,2)"`
}

func (Access) TableName() string {
	return "access"
}

// store node info to db
func (fm *Access) CreateFileMemo() error {
	// store memodid
	return G_DB.Create(fm).Error
}
