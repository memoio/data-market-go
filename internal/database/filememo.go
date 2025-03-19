package database

type FileMemo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	FileDID     string `gorm:"column:file_did;type:TEXT;not null;primaryKey"`
	MemoDID     string `gorm:"column:memo_did;type:TEXT;not null;primaryKey"`
	UserAddress string `gorm:"column:user_address;type:TEXT;not null"`
	AddType     int    `gorm:"column:add_type;type:INTEGER;not null;check:add_type IN (1,2)"`
}

func (FileMemo) TableName() string {
	return "file_memo"
}

// store node info to db
func (fm *FileMemo) CreateFileMemo() error {
	// store memodid
	return G_DB.Create(fm).Error
}
